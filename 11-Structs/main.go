package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
	End Endereco
}

type Usuario struct {
	Nome  string
	Idade int
	Ativo bool
	End   Endereco
}

func main() {
	evelyn := Cliente{
		Nome:  "Evelyn",
		Idade: 30,
		Ativo: true,
	}

	evelyn.Ativo = false
	evelyn.Cidade = "São Paulo"
	evelyn.Endereco.Cidade = "São Paulo"
	evelyn.End.Logradouro = "Rua I"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", evelyn.Nome, evelyn.Idade, evelyn.Ativo)
}
