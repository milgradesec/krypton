package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/sys/windows/registry"
)

// Client almacena la información que enviará al servidor
// como telemetría en las actualizaciones
type Client struct {
	ID      string `json:"id,omitempty"`
	Version string `json:"version,omitempty"`
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

func serverNotify() {
	c := Client{ID: getID(), Version: version}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(c)
	_, err := http.Post("https://paesacybersecurity.eu/api/telemetry/new", "application/json", b)
	if err != nil {
		log.Println(err)
	}
}
