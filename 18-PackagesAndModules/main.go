package main

import (
	"fmt"

	"github.com/EvyOliveira/goExpert/18-PackagesAndModules/mathematics"
	"github.com/google/uuid" 
)

func main() {
	soma := mathematics.Soma(10, 20)

	carro := mathematics.Carro{Marca: "Fiat"}
	fmt.Println(carro.Andar())
	fmt.Println(mathematics.A)
	fmt.Println(" Resultado: $v", soma)
	fmt.Println(uuid.New())
}
