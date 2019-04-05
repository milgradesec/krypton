package main

import (
	"fmt"
	"os"
)

const (
	version = "0.4"
)

func main() {
	parseCommandLine()
}

func showHelp() {
	fmt.Println("")
	fmt.Println("Krypton " + version + "  --  Security Configuration Updater")
	fmt.Println("")
	fmt.Println("Uso: Krypton <comando>")
	fmt.Println("")
	fmt.Println("Comandos:")
	fmt.Println("	--help, -h")
	fmt.Println("	--install")
	fmt.Println("	--update")
	fmt.Println("	--force-update")
	fmt.Println("	--upgrade")
}

func parseCommandLine() {
	args := os.Args
	if len(args) == 1 {
		showHelp()
	} else if args[1] == "-h" || args[1] == "--help" {
		showHelp()
	} else if args[1] == "-i" || args[1] == "--install" {
		install()
	} else if args[1] == "--update" {
		updateConfiguration(false)
	} else if args[1] == "--force-update" {
		updateConfiguration(true)
	} else if args[1] == "--upgrade" {
		update()
	} else {
		showHelp()
		fmt.Printf("Error: comando no válido: %s", args[1])
	}
}
