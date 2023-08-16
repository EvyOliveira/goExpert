package main

import "fmt"

func main() {
	salarios := map[string]int{"Evelyn": 5000, "João": 6000, "Maria": 3500}
	fmt.Println(salarios["Evelyn"])
	delete(salarios, "Evelyn")

	salarios["Evy"] = 3000
	fmt.Println(salarios["Evy"])
}
