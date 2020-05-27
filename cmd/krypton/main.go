package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/milgradesec/krypton/internal/installer"
	"github.com/milgradesec/krypton/internal/system"
	"github.com/milgradesec/krypton/internal/updater"
)

func main() {
	var (
		versionFlag     = flag.Bool("version", false, "Show version information")
		installFlag     = flag.Bool("install", false, "Instala Krypton en el sistema")
		updateFlag      = flag.Bool("update", false, "Actualiza la configuración de seguridad si hay cambios")
		forceUpdateFlag = flag.Bool("force-update", false, "Actualiza la configuración de seguridad")
		upgradeFlag     = flag.Bool("upgrade", false, "Actualiza Krypton a la ultima versión")
		helpFlag        = flag.Bool("help", false, "Muestra los comandos disponibles")
	)
	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Println("extra command line arguments.")
		os.Exit(1)
	}

	if *versionFlag {
		fmt.Println("Krypton " + Version)
		fmt.Printf("%s/%s, %s, %s\n", runtime.GOOS, runtime.GOARCH, runtime.Version(), Version)
	}

	if *installFlag {
		err := installer.Install()
		if err != nil {
			fmt.Printf("Error instalando Krypton: %v\n", err)
		}
	}

	if *updateFlag {
		err := system.UpdateConfig(false)
		if err != nil {
			fmt.Printf("Error actualizando configuración: %v\n", err)
		}
	}

	if *forceUpdateFlag {
		err := system.UpdateConfig(true)
		if err != nil {
			fmt.Printf("Error actualizando configuración: %v\n", err)
		}
	}

	if *upgradeFlag {
		err := updater.Update(Version)
		if err != nil {
			fmt.Printf("Error actualizando Krypton: %v\n", err)
		}
	}

	if *helpFlag {
		flag.PrintDefaults()
	}
}

var (
	Version string
)
