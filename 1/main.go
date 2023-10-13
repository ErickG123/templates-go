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
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
