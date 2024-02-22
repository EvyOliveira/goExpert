package main

func Soma(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"Evelyn": 10000, "Stephen": 3000, "Bob": 4000}
	println(Soma(m))
}
