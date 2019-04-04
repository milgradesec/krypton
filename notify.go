package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/sys/windows/registry"
)

// ClientTelemetry almacena la información que enviará al servidor
// como telemetría durante las actualizaciones
type ClientTelemetry struct {
	ID      string `json:"id,omitempty"`
	Version string `json:"version,omitempty"`
	OSBuild string `json:"osbuild,omitempty"`
}

func uploadTelemetry() {
	c := ClientTelemetry{
		ID:      getID(),
		OSBuild: getWindowsVersion(),
		Version: version,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(c)

	_, err := http.Post("https://paesacybersecurity.eu/api/telemetry/new", "application/json", b)
	if err != nil {
		log.Println(err)
	}
}

func createNewID() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Krypton", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	id := uuid.New()
	idString := id.String()
	k.SetStringValue("uuid", idString)
	return idString
}

func getID() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Krypton", registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	id, _, err := k.GetStringValue("uuid")
	if err != nil {
		return createNewID()
	}
	return id
}
