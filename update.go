package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

// update actualiza Krypton a la versión más reciente disponible en el servidor
func update() error {
	resp, err := http.Get("https://dl.paesacybersecurity.eu/krypton/stable/krypton.version")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("update: response status != 200")
	}

	v, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if string(v) == version {
		log.Println("Krypton esta actualizado.")
		return nil
	}

	url := "https://dl.paesacybersecurity.eu/krypton/stable/Krypton.exe"
	path := "C:/Program Files/Krypton/Updates/Krypton.exe"
	err = downloadToFile(url, path)
	if err != nil {
		return err
	}

	cmd := exec.Command(path, "--install")
	err = cmd.Start()
	if err != nil {
		return err
	}
	return nil
}
