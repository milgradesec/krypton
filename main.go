package main

import (
	"flag"
	"log"
)

const (
	version = "0.5"
)

var (
	installFlag     = flag.Bool("install", false, "Instala Krypton en el sistema")
	updateFlag      = flag.Bool("update", false, "Actualiza la configuración de seguridad si hay cambios")
	forceUpdateFlag = flag.Bool("force-update", false, "Actualiza la configuración de seguridad")
	upgradeFlag     = flag.Bool("upgrade", false, "Actualiza Krypton a la ultima versión")
	helpFlag        = flag.Bool("help", false, "Muestra los comandos disponibles")
)

func main() {
	log.Println("Krypton " + version + "  --  Security Configuration Updater")
	flag.Parse()
	if *installFlag {
		err := installKrypton()
		if err != nil {
			log.Printf("Error instalando Krypton: %v\n", err)
		}
	} else if *updateFlag {
		updateConfiguration(false)
	} else if *forceUpdateFlag {
		updateConfiguration(true)
	} else if *upgradeFlag {
		err := update()
		if err != nil {
			log.Printf("Error actualizando Krypton: %v\n", err)
		}

	} else if *helpFlag {
		flag.PrintDefaults()
	} else {
		flag.PrintDefaults()
	}
}
