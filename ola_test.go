package main

import "testing"

func TestOla(t *testing.T) {

	//subtest
	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Luc")
		esperado := "Olá, Luc"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	//subtest
	t.Run("diz 'Olá, mundo', quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá, mundo"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

}
