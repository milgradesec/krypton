package main

import (
	"flag"
	"fmt"
)

const (
	version = "0.4.2"
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
		err := installKrypton()
		if err != nil {
			fmt.Printf("Error instalando: %v\n", err)
		}
	} else if *updateFlag {

	} else if *forceUpdateFlag {

	} else if *upgradeFlag {

	} else if *helpFlag {
		flag.PrintDefaults()
	} else {
		flag.PrintDefaults()
	}
}
