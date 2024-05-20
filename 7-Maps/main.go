package main

import "fmt"

func main() {
	salarios := map[string]int{"Evelyn": 5000, "João": 6000, "Maria": 3500}
	delete(salarios, "Evelyn")

	salarios["Evy"] = 3000

	//sal := make(map[string]int)
	//sal1 := map[string]int{}
	//sal1["Evelyn"] = 10000

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é de %d\n", salario)
	}
}
