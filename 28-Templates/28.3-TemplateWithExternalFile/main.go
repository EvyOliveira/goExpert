package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"Golang", 40},
		{"Python", 30},
		{"Java", 50},
	})
	if err != nil {
		panic(err)
	}
}
