package main

type MyNumber int
type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Evelyn": 10000, "Stephen": 3000, "Bob": 4000}
	m2 := map[string]float64{"Evelyn": 10000.30, "Stephen": 3000.5, "Bob": 4000.0}
	m3 := map[string]MyNumber{"Evelyn": 10000, "Stephen": 3000, "Bob": 4000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 11))
}
