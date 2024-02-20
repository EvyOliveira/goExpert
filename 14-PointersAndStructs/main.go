package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Evelyn Oliveira"
	fmt.Printf("A cliente %v andou\n", c.nome)
}

func main() {
	evelyn := Cliente{
		nome: "Evelyn",
	}
	evelyn.andou()
	fmt.Printf("O valor da struct com nome %v", evelyn.nome)
}
