package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

// Actualiza Krypton a la versión mas reciente disponible
// en el servidor de actualizaciones
func update() {
	UploadTelemetry()
	if isUpdateAvailable() {
		log.Println("Hay nueva versión disponible")

		path := "C:/Program Files/Krypton/Updates/Krypton.exe"
		currentChannel := LoadCurrentChannel()
		err := downloadToFile(currentChannel.updateURL, path)
		if err != nil {
			log.Fatal("Error al descargar actualización")
		}
		log.Println("Actualización descargada correctamente")

		// Ejecutar la nueva versión para que se instale
		cmd := exec.Command(path, "--install")
		err = cmd.Start()
	} else {
		log.Println("No hay nueva versión disponible")
	}

}

// Comprueba si hay una nueva versión de Krypton disponible
// comparando su versión con la que contiene el archivo krypton.version
// descargado del servidor de actualizaciones
func isUpdateAvailable() bool {
	currentChannel := LoadCurrentChannel()
	resp, err := http.Get(currentChannel.updateVersionURL)
	if err != nil {
		return false
	}
	if resp.StatusCode != 200 {
		return false
	}
	defer resp.Body.Close()

	newVersion, err := ioutil.ReadAll(resp.Body)
	if string(newVersion) == version {
		return false
	}
	return true
}
