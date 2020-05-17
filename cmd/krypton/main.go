package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"github.com/milgradesec/krypton/internal/installer"
)

func main() {
	fmt.Printf("KRYPTON-%s\n", Version)

	var (
		versionFlag     = flag.Bool("version", false, "Show version information.")
		installFlag     = flag.Bool("install", false, "Instala Krypton en el sistema")
		updateFlag      = flag.Bool("update", false, "Actualiza la configuración de seguridad si hay cambios")
		forceUpdateFlag = flag.Bool("force-update", false, "Actualiza la configuración de seguridad")
		upgradeFlag     = flag.Bool("upgrade", false, "Actualiza Krypton a la ultima versión")
		helpFlag        = flag.Bool("help", false, "Muestra los comandos disponibles")
	)
	flag.Parse()

	if *helpFlag {
		flag.PrintDefaults()
		return
	}

	if *versionFlag {
		fmt.Printf("%s/%s, %s, %s\n", runtime.GOOS, runtime.GOARCH, runtime.Version(), Version)
		return
	}

	if *installFlag {
		if err := installer.Install(); err != nil {
			log.Fatal(err)
		}
	}

	if *updateFlag || *forceUpdateFlag {

	}

	if *upgradeFlag {

	}

}

var (
	Version string
)
