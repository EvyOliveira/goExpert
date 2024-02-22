package main

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"Evelyn": 10000, "Stephen": 3000, "Bob": 4000}
	m2 := map[string]float64{"Evelyn": 10000.30, "Stephen": 3000.5, "Bob": 4000.0}
	println(SomaInteiro(m))
	println(SomaFloat(m2))
}
