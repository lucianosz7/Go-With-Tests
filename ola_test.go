package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		// t.Helper -> Diz que é um método auxiliar e não será relatado no teste
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	//subtest
	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Luc", "")
		esperado := "Olá, Luc"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	//subtest
	t.Run("diz 'Olá, mundo', quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T) {
		resultado := Ola("Elodie", "francês")
		esperado := "Bonjour, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

}
