package main

import "testing"

func TestGetID(t *testing.T) {
	id := getID()
	if id != "e19ef7ef-0460-42c6-b6f3-b200771149fd" {
		t.Error("ID no coincide")
	}
}
