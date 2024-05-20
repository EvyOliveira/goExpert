package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Evelyn"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var array [3]int
	array[0] = 10
	array[1] = 20
	array[2] = 30

	for i, v := range array {
		fmt.Printf(" O valor do indice %d é %d\n", i, v)
	}

	fmt.Println(len(array) - 1)
	fmt.Println(len(array))
	fmt.Println(array[len(array)-1])

	fmt.Printf("O tipo de E é %T", e)
	fmt.Printf("O tipo de E é %v", e)

	fmt.Printf("O tipo de E é %T", f)
	fmt.Printf("O tipo de E é %v", f)
}

func x() {
}
