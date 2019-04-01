package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func update() {
	if isUpdateAvailable() {
		path := "C:\\Program Files\\Krypton\\Updates\\Krypton.exe"
		url := "https://paesacybersecurity.eu/krypton/Krypton.exe"
		err := downloadToFile(url, path)
		if err != nil {
			log.Println("Error al descargar actualizacion")
		} else {
			log.Println("Actualizacion descargada correctamente")
			cmd := exec.Command(path, "--install")
			err = cmd.Start()
		}
	}
}

func isUpdateAvailable() bool {
	resp, err := http.Get("https://paesacybersecurity.eu/krypton/krypton.version")
	if err != nil || resp.StatusCode != 200 {
		log.Println("Error al comprobar si hay nueva version")
		return false
	}
	newVersion, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if string(newVersion) == version {
		log.Println("Krypton esta actualizado")
		return false
	}

	log.Println("Hay nueva version disponible")
	return true
}
