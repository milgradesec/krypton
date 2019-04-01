package main

import (
	"os"
	"os/exec"
	"time"

	"golang.org/x/sys/windows/registry"
)

func install() {

	// Esperar...
	time.Sleep(15 * time.Second)

	if isKryptonInstalled() {
		// Eliminar archivos y carpetas del antiguo Krypton
		os.Remove("C:\\Program Files\\Krypton\\7z.exe")
		os.Remove("C:\\Program Files\\Krypton\\data.zip")
		os.Remove("C:\\Program Files\\Krypton\\update.zip")
		os.Remove("C:\\Program Files\\Krypton\\Settings.xml")
		os.RemoveAll("C:\\Program Files\\Krypton\\data")
		os.RemoveAll("C:\\Program Files\\Krypton\\update")
		registry.DeleteKey(registry.LOCAL_MACHINE, "SOFTWARE\\Krypton")
	} else {
		os.Mkdir("C:\\Program Files\\Krypton", os.ModeDir)
	}

	// Crear nuevas carpetas
	os.Mkdir("C:\\Program Files\\Krypton\\Updates", os.ModeDir)
	os.Mkdir("C:\\Program Files\\Krypton\\Logs", os.ModeDir)

	// Mover ejecutable actual a la carpeta de instalación
	path, _ := os.Executable()
	copyFile(path, "C:\\Program Files\\Krypton\\Krypton.exe")

	registry.CreateKey(registry.LOCAL_MACHINE, "SOFTWARE\\Krypton", registry.ALL_ACCESS)
	createScheduledTasks()
}

// Comprueba si existe la carpeta de instalación de Krypton
func isKryptonInstalled() bool {
	_, err := os.Stat("C:\\Program Files\\Krypton")
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Crea las tareas programadas para ejecutarse periódicamente
func createScheduledTasks() {
	path := "C:\\Program Files\\Krypton\\Krypton.exe"
	cmd := exec.Command("schtasks.exe", "/Create", "/SC", "HOURLY", "/TN", "KryptonUpdate",
		"/RU", "SYSTEM", "/F", "/TR", path+" --update")
	cmd.Run()

	cmd = exec.Command("schtasks.exe", "/Create", "/SC", "DAILY", "/TN", "KryptonUpgrade",
		"/RU", "SYSTEM", "/F", "/TR", path+" --upgrade")
	cmd.Run()
}
