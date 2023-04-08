package main

import "fmt"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return prefixoSaudacao(idioma) + nome
}

func main() {
	fmt.Println(Ola("Mundo", ""))
}

func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case "espanhol":
		prefixo = prefixoOlaEspanhol
	case "francês":
		prefixo = prefixoOlaFrances
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}
