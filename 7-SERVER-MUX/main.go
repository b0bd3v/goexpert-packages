package main

import "net/http"

// Multiplexer (Mux) is a router that matches the incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions.
// https://pkg.go.dev/net/http#ServeMux

func main() {
	finish := make(chan bool)

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "Blog Page"})

	go func() {
		http.ListenAndServe(":8080", mux)
	}()

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	go func() {
		http.ListenAndServe(":8082", mux2)
	}()

	<-finish
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(b.title))
}
