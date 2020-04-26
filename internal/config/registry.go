package config

import (
	"log"

	"golang.org/x/sys/windows/registry"
)

func getLastUpdateHash() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		"SOFTWARE\\Krypton", registry.QUERY_VALUE)
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

func setLastUpdateHash(hash string) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		"SOFTWARE\\Krypton", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	k.SetStringValue("lastUpdateHash", hash)
}

func getWindowsVersion() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		"SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", registry.QUERY_VALUE)
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

func setLastUpdateWindowsVersion(buildNumber string) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		"SOFTWARE\\Krypton", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	k.SetStringValue("lastBuildNumber", buildNumber)
}

func getWindowsPatchNumber() uint64 {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		"SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion", registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	patchNumber, _, err := k.GetIntegerValue("UBR")
	if err != nil {
		log.Fatal(err)
	}
	return patchNumber
}

func getLastUpdateWindowsVersion() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		"SOFTWARE\\Krypton", registry.QUERY_VALUE)
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
