package main

import (
	"log"

	"golang.org/x/sys/windows/registry"
)

// Canales de actualización de Krypton
const (
	// STABLE recibe las versiones estables (predeterminado)
	UpdateChannelStable = 0
	// BETA recibe las versiones en fase de pruebas
	UpdateChannelBeta = 1
	// DEV recibe las versiones en fase de desarrollo (inestable)
	UpdateChannelDev = 2
)

// Canales de actualización de configuracion
const (
	// STABLE instala las configuraciones estables
	ConfigChannelStable = 0
	// TEST instala las configuraciones de prueba
	ConfigChannelTest = 1
)

// UpdateChannels almacena los valores de los canales de actualización
// de Krypton y la configuración de cada uno
type UpdateChannels struct {
	updateChannel         int
	configChanel          int
	updateVersionURL      string
	updateURL             string
	configurationURL      string
	exploitMitigationsURL string
}

func loadCurrentChannel() *UpdateChannels {
	c := UpdateChannels{}
	url := "https://paesacybersecurity.eu/krypton/"
	var dir string

	c.updateChannel = getCurrentUpdateChannel()
	switch c.updateChannel {
	case UpdateChannelStable:
		dir = "stable"
	case UpdateChannelBeta:
		dir = "beta"
	case UpdateChannelDev:
		dir = "dev"
	}
	c.updateVersionURL = url + dir + "/krypton.version"
	c.updateURL = url + dir + "/Krypton.exe"

	c.configChanel = getCurrentConfigChannel()
	switch c.configChanel {
	case ConfigChannelStable:
		dir = "config/stable"
	case ConfigChannelTest:
		dir = "config/test"
	}
	c.configurationURL = url + dir + "/config.zip"
	c.exploitMitigationsURL = url + dir + "/Settings.xml"
	return &c
}

func getCurrentUpdateChannel() int {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Krypton", registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	channel, _, err := k.GetIntegerValue("updateChannel")
	if err != nil {
		// Si no se especifica ninguno se utiliza el estable
		return UpdateChannelStable
	}
	return int(channel)
}

func getCurrentConfigChannel() int {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Krypton", registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	channel, _, err := k.GetIntegerValue("configChannel")
	if err != nil {
		// Si no se especifica ninguno se utiliza el estable
		return ConfigChannelStable
	}
	return int(channel)
}
