package main

import (
	"log"
	"net/http"
	"text/template"
)

type Curso struct {
	Title        string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))

		err := t.Execute(w, Cursos{
			{Title: "Go", CargaHoraria: 40},
			{Title: "Python", CargaHoraria: 35},
			{Title: "Java", CargaHoraria: 30},
			{Title: "C#", CargaHoraria: 25},
		})

		if err != nil {
			log.Fatal(err)
		}

	})

	http.ListenAndServe(":8080", nil)
}
