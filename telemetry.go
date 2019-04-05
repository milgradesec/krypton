package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"golang.org/x/sys/windows/registry"
)

// ClientTelemetry almacena la información que enviará al servidor
// como telemetría durante las actualizaciones
type ClientTelemetry struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	OSBuild string `json:"osbuild"`
}

// UploadTelemetry envía la información de telemetría al servidor
func UploadTelemetry() {
	c := ClientTelemetry{
		ID:      GetComputerID(),
		Version: version,
		OSBuild: getWindowsVersion() + "." + strconv.FormatUint(getWindowsPatchNumber(), 10),
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(c)

	_, err := http.Post("https://paesacybersecurity.eu/api/telemetry/new",
		"application/json", b)
	if err != nil {
		log.Println(err)
	}
}

// CreateNewID genera un Identificador Único y lo guarda en el registro
func CreateNewID() string {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE,
		"SOFTWARE\\Krypton", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	id := uuid.New()
	idString := id.String()
	key.SetStringValue("uuid", idString)
	return idString
}

// GetComputerID devuelve el ID almacenado en el registro
func GetComputerID() string {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE,
		"SOFTWARE\\Krypton", registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	id, _, err := key.GetStringValue("uuid")
	if err != nil {
		return CreateNewID()
	}
	return id
}
