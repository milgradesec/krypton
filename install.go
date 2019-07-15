package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"golang.org/x/sys/windows/registry"
)

const (
	updateDelay      = 5 * time.Second
	kryptonDirectory = "C:/Program Files/Krypton/"
)

func isAlreadyInstalled() bool {
	_, err := os.Stat(kryptonDirectory)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func cleanup() {
	os.Remove(kryptonDirectory + "7z.exe")
	os.Remove(kryptonDirectory + "data.zip")
	os.Remove(kryptonDirectory + "update.zip")
	os.Remove(kryptonDirectory + "Settings.xml")
	os.RemoveAll(kryptonDirectory + "data")
	os.RemoveAll(kryptonDirectory + "update")
}

func createRegistryKeys() error {
	_, _, err := registry.CreateKey(registry.LOCAL_MACHINE,
		`SOFTWARE\\Krypton`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	return nil
}

func createScheduledTasks() error {
	path := kryptonDirectory + "/Krypton.exe"
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
		err := os.Mkdir(kryptonDirectory, os.ModeDir)
		if err != nil {
			return err
		}
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
	err = copyFile(path, kryptonDirectory+"Krypton.exe")
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
