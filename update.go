package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func update() {
	if isUpdateAvailable() {
		log.Println("Hay nueva versión disponible")

		path := "C:\\Program Files\\Krypton\\Updates\\Krypton.exe"
		url := "https://paesacybersecurity.eu/krypton/Krypton.exe"
		err := downloadToFile(url, path)
		if err != nil {
			log.Fatal("Error al descargar actualización")
		}
		log.Println("Actualización descargada correctamente")
		cmd := exec.Command(path, "--install")
		err = cmd.Start()

	} else {
		log.Println("No hay nueva versión disponible")
	}

}

func isUpdateAvailable() bool {
	resp, err := http.Get("https://paesacybersecurity.eu/krypton/krypton.version")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	newVersion, err := ioutil.ReadAll(resp.Body)
	if string(newVersion) == version {
		return false
	}
	return true
}
