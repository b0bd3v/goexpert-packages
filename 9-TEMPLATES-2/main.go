package main

import (
	"log"
	"os"
	"text/template"
)

type Curso struct {
	Title        string
	CargaHoraria int
}

func main() {
	curso := Curso{Title: "Curso de Go", CargaHoraria: 20}

	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{ .Title }}\nCarga Horária: {{ .CargaHoraria }} horas"))

	err := t.Execute(os.Stdout, curso)

	if err != nil {
		log.Fatal(err)
	}
}
