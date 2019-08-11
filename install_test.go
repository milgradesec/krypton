package main

import (
	"os"
	"testing"
)

// schtasks.exe /query /tn KryptonUpdate /xml

func Test_installKrypton(t *testing.T) {
	path, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}

	hash1 := computeFileSHA1(path)
	if hash1 == "" {
		t.Fatal("Hash vacío")
	}

	err = install()
	if err != nil {
		t.Fatal(err)
	}

	hash2 := computeFileSHA1("C:/Program Files/Krypton/Krypton.exe")
	if hash2 == "" {
		t.Fatal("Hash vacío")
	}

	if hash1 != hash2 {
		t.Fatal("No coinciden las firmas de los archivos")
	}
}
