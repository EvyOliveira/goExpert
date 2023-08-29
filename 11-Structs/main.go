package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	evelyn := Cliente{
		Nome:  "Evelyn",
		Idade: 30,
		Ativo: true,
	}
	evelyn.Ativo = false
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", evelyn.Nome, evelyn.Idade, evelyn.Ativo)
}
