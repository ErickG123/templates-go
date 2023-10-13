package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	// Templates
	curso := Curso{"Go", 40}

	// Criando um Template
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))

	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
