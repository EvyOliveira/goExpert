package main

func main() {
	a := 10
	println(&a)

	var ponteiro *int = &a
	println(ponteiro)

	*ponteiro = 20
	println(a)

	b := &a
	println(*b)

	*b = 30
	println(a)
	println(*b)
}
