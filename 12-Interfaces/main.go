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

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func (e Empresa) Desativar() {

}

func main() {
	evelyn := Cliente{
		Nome:  "Evelyn",
		Idade: 30,
		Ativo: true,
	}

	minhaEmpresa := Empresa{}
	Desativacao(minhaEmpresa)

	evelyn.Cidade = "São Paulo"
	evelyn.Endereco.Cidade = "São Paulo"
	evelyn.End.Logradouro = "Rua I"
	evelyn.Ativo = false
	evelyn.Desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", evelyn.Nome, evelyn.Idade, evelyn.Ativo)
}
