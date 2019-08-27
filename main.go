package main

import (
	"flag"
	"fmt"
)

const (
	version = "1.0"
)

var (
	installFlag     = flag.Bool("install", false, "Instala Krypton en el sistema")
	updateFlag      = flag.Bool("update", false, "Actualiza la configuración de seguridad si hay cambios")
	forceUpdateFlag = flag.Bool("force-update", false, "Actualiza la configuración de seguridad")
	upgradeFlag     = flag.Bool("upgrade", false, "Actualiza Krypton a la ultima versión")
	helpFlag        = flag.Bool("help", false, "Muestra los comandos disponibles")
)

func main() {
	fmt.Println("Krypton " + version + "  --  Security Configuration Updater")
	flag.Parse()
	if *installFlag {
		err := install()
		if err != nil {
			fmt.Printf("Error instalando Krypton: %v\n", err)
		}
	} else if *updateFlag {
		err := updateConfig(false)
		if err != nil {
			fmt.Printf("Error actualizando configuración: %v\n", err)
		}

	} else if *forceUpdateFlag {
		err := updateConfig(true)
		if err != nil {
			fmt.Printf("Error actualizando configuración: %v\n", err)
		}

	} else if *upgradeFlag {
		err := update()
		if err != nil {
			fmt.Printf("Error actualizando Krypton: %v\n", err)
		}

	} else if *helpFlag {
		flag.PrintDefaults()
	} else {
		flag.PrintDefaults()
	}
}
