package main

import (
	"log"

	"golang.org/x/sys/windows/registry"
)

const (
	channelStable = 0
	channelTest   = 1
)

// UpdateChannel almacena los valores de los canales de actualización
// de Krypton y la configuración de cada uno
type UpdateChannel struct {
	UpdateVersionURL      string
	UpdateURL             string
	ConfigurationURL      string
	ExploitMitigationsURL string
}

func loadCurrentChannel() *UpdateChannel {
	c := UpdateChannel{}
	url := "https://paesacybersecurity.eu/krypton/"
	var dir string

	switch getCurrentUpdateChannel() {
	case channelStable:
		dir = "stable"
	case channelTest:
		dir = "dev"
	}
	c.UpdateVersionURL = url + dir + "/krypton.version"
	c.UpdateURL = url + dir + "/Krypton.exe"

	switch getCurrentConfigChannel() {
	case channelStable:
		dir = "config/stable"
	case channelTest:
		dir = "config/test"
	}
	c.ConfigurationURL = url + dir + "/config.zip"
	c.ExploitMitigationsURL = url + dir + "/Settings.xml"
	return &c
}

func getCurrentUpdateChannel() int {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		"SOFTWARE\\Krypton", registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	channel, _, err := k.GetIntegerValue("updateChannel")
	if err != nil {
		// Si no se especifica ninguno se utiliza el estable
		return channelStable
	}
	return int(channel)
}

func getCurrentConfigChannel() int {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		"SOFTWARE\\Krypton", registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	channel, _, err := k.GetIntegerValue("configChannel")
	if err != nil {
		// Si no se especifica ninguno se utiliza el estable
		return channelStable
	}
	return int(channel)
}
