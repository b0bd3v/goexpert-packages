package main

import "os"

type Curso struct {
	Title        string
	CargaHoraria int
}

func main() {
	curso := Curso{Title: "Curso de Go", CargaHoraria: 20}

	tmpl := template.New("CursoTemplate")

	tmpl, _ = tmpl.Parse("Curso: {{ .Title }}\nCarga Horária: {{ .CargaHoraria }} horas")

	err := tmpl.Execute(os.Stdout, curso)
}
