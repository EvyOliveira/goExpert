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
	fmt.Printf("O tipo de E é %T", e)
	fmt.Printf("O tipo de E é %v", e)

	fmt.Printf("O tipo de E é %T", f)
	fmt.Printf("O tipo de E é %v", f)
}

func x() {
}
