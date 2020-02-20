package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

// update actualiza Krypton a la versión más reciente disponible en el servidor
func update() error {
	resp, err := http.Get("https://dl.paesacybersecurity.eu/krypton/krypton.version")
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
		fmt.Println("Krypton esta actualizado.")
		return nil
	}

	url := "https://dl.paesacybersecurity.eu/krypton/krypton.exe"
	path := "C:/Program Files/Krypton/Updates/krypton.exe"

	os.Mkdir("C:/Program Files/Krypton/Updates", os.ModeDir)
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
