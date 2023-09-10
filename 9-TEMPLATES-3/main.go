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

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Cursos{
		{Title: "Go", CargaHoraria: 40},
		{Title: "Python", CargaHoraria: 35},
		{Title: "Java", CargaHoraria: 30},
		{Title: "C#", CargaHoraria: 25},
	})

	if err != nil {
		log.Fatal(err)
	}
}
