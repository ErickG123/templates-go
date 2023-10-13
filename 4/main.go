package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	// Adicionando uma Rota
	http.HandleFunc("/", HomeHandler)
	// Criando um Web Server
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Criando um Template
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	// Estou dando um Execute no meu ResponseWriter
	// Ele pega o Response e joga para o Write que vai gerar o Response
	err := t.Execute(w, Cursos{
		{"Go", 40},
		{"Python", 30},
	})

	if err != nil {
		panic(err)
	}
}
