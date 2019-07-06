package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func newVersionAvailable() bool {
	currentChannel := loadCurrentChannel()
	resp, err := http.Get(currentChannel.updateVersionURL)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	newVersion, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	if string(newVersion) == version {
		return false
	}
	return true
}

func update() {
	if newVersionAvailable() {
		log.Println("Hay nueva versión disponible")

		path := "C:/Program Files/Krypton/Updates/Krypton.exe"
		currentChannel := loadCurrentChannel()
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
