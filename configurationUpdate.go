package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"golang.org/x/sys/windows/registry"
)

func updateConfiguration(force bool) {

	// Always update exploit mitigations
	updateExploitMitigations()

	path := "C:/Program Files/Krypton/Updates/config.zip"
	currentChannel := loadChannelsInfo()
	err := downloadToFile(currentChannel.configurationURL, path)
	if err != nil {
		log.Println("Error al descargar la configuracion de seguridad")
		log.Fatal(err)
	}
	if force == false {
		newHash := getFileHash(path)
		oldHash := getUpdateHash()
		if newHash == oldHash {
			log.Println("No hay cambios de configuracion")
			os.Exit(0)
		}
		log.Println("Hay nueva configuracion disponible")
		saveUpdateHash(newHash)
	}

	os.RemoveAll("C:\\Program Files\\Krypton\\Updates\\config")
	err = unzip(path, "C:\\Program Files\\Krypton\\Updates")
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir("C:/Program Files/Krypton/Updates/config")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		runPowershellScript("./"+f.Name(), "C:/Program Files/Krypton/Updates/config")
	}
}

func getUpdateHash() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Krypton", registry.QUERY_VALUE)
	if err != nil {
		return ""
	}
	defer k.Close()

	hash, _, err := k.GetStringValue("lastUpdateHash")
	if err != nil {
		return ""
	}
	return hash
}

func saveUpdateHash(hash string) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Krypton", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	k.SetStringValue("lastUpdateHash", hash)
}

func getWindowsVersion() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	buildNumber, _, err := k.GetStringValue("CurrentBuildNumber")
	if err != nil {
		log.Fatal(err)
	}
	return buildNumber
}

func getLastUpdateWindowsVersion() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Krypton", registry.QUERY_VALUE)
	if err != nil {
		return ""
	}
	defer k.Close()

	buildNumber, _, err := k.GetStringValue("lastBuildNumber")
	if err != nil {
		return ""
	}
	return buildNumber
}

func updateExploitMitigations() {
	url := "https://paesacybersecurity.eu/krypton/Settings.xml"
	path := "C:/Program Files/Krypton/Updates/Settings.xml"
	err := downloadToFile(url, path)
	if err != nil {
		log.Println("Error al descargar la configuracion contra exploits")
		return
	}
	runPowershellScript("Set-ProcessMitigation -PolicyFilePath Settings.xml",
		"C:\\Program Files\\Krypton\\Updates")
	log.Println("Actualizada configuracion contra exploits")
}

func runPowershellScript(flags string, workingDir string) {
	cmd := exec.Command("powershell.exe", "-ExecutionPolicy", "Bypass", flags)
	cmd.Dir = workingDir
	cmd.Start()
}
