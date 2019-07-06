package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"golang.org/x/sys/windows/registry"
)

const (
	updateDelay = 5 * time.Second
	kryptonDir  = "C:/Program Files/Krypton"
)

func isAlreadyInstalled() bool {
	_, err := os.Stat("C:/Program Files/Krypton")
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func cleanup() {
	os.Remove(kryptonDir + "/7z.exe")
	os.Remove(kryptonDir + "/data.zip")
	os.Remove(kryptonDir + "/update.zip")
	os.Remove(kryptonDir + "/Settings.xml")
	os.RemoveAll(kryptonDir + "/data")
	os.RemoveAll(kryptonDir + "/update")
	os.RemoveAll(kryptonDir + "/Logs")
}

func createRegistryKeys() error {
	registry.CreateKey(registry.LOCAL_MACHINE,
		"SOFTWARE/Krypton", registry.ALL_ACCESS)
	return nil
}

func createScheduledTasks() error {
	path := kryptonDir + "/Krypton.exe"
	cmd := exec.Command("schtasks.exe", "/Create", "/SC", "HOURLY", "/TN",
		"KryptonUpdate", "/RU", "SYSTEM", "/F", "/TR", path+" --update")
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("schtasks.exe", "/Create", "/SC", "DAILY", "/TN",
		"KryptonUpgrade", "/RU", "SYSTEM", "/F", "/TR", path+" --upgrade")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func installKrypton() error {
	if !isAlreadyInstalled() {
		err := os.Mkdir(kryptonDir, os.ModeDir)
		if err != nil {
			return err
		}
		fmt.Println("Krypton no esta instalado.")
	}
	fmt.Println("Instalando...")

	// Esperar para evitar errores
	time.Sleep(updateDelay)
	cleanup()

	// Mover ejecutable actual a la carpeta de instalación
	path, err := os.Executable()
	if err != nil {
		return err
	}
	err = copyFile(path, kryptonDir+"/Krypton.exe")
	if err != nil {
		return err
	}

	err = createRegistryKeys()
	if err != nil {
		return err
	}
	err = createScheduledTasks()
	if err != nil {
		return err
	}
	fmt.Println("Instalado correctamente.")
	return nil
}
