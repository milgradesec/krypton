package main

import (
	"fmt"
	"io/ioutil"
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

func update() error {
	if !newVersionAvailable() {
		fmt.Println("No hay nueva versión disponible.")
	} else {
		fmt.Println("Hay nueva versión disponible.")
		currentChannel := loadCurrentChannel()
		path := kryptonDir + "/Updates/Krypton.exe"
		err := downloadToFile(currentChannel.updateURL, path)
		if err != nil {
			return err
		}

		cmd := exec.Command(path, "--install")
		err = cmd.Start()
		if err != nil {
			return err
		}
	}
	return nil
}
