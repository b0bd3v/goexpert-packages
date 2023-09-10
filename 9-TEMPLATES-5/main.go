package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Title        string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))

	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Python", 40},
		{"C++", 40},
	})

	if err != nil {
		panic(err)
	}
}
