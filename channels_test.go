package main

import "testing"

func TestGetCurrentChannel(t *testing.T) {
	uchannel := getCurrentUpdateChannel()
	switch uchannel {
	case 0:
		t.Log("Canal de actualización estable")
	case 1:
		t.Log("Canal de actualización beta")
	case 2:
		t.Log("Canal de actualización desarrollo")
	default:
		t.Fatalf("Canal de actualización erroneo")
	}

	cchannel := getCurrentConfigChannel()
	switch cchannel {
	case 0:
		t.Log("Canal de configuración estable")
	case 1:
		t.Log("Canal de configuración inestable")
	default:
		t.Fatalf("Canal de configuración erroneo")
	}
}
