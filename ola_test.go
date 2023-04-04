package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Luc")
	esperado := "Olá, Luc"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
