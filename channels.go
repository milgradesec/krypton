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

func getCurrentUpdateChannel() int {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Krypton", registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	channel, _, err := k.GetIntegerValue("updateChannel")
	if err != nil {
		// Si no se especifica ninguno se utiliza en estable
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
		// Si no se especifica ninguno se utiliza en estable
		return ConfigChannelStable
	}
	return int(channel)
}
