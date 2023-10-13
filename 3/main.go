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
	// Criando um Template
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Python", 30},
	})

	if err != nil {
		panic(err)
	}
}
