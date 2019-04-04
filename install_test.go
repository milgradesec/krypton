package main

import "testing"

func TestIsKryptonInstalled(t *testing.T) {
	if !isKryptonInstalled() {
		t.Fatal("Krypton no esta instalado")
	}
	t.Log("Krypton esta instalado")
}
