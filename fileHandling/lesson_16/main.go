package main

import (
	"log"
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleTemplate)
	log.Fatal(http.ListenAndServe(":8080", mux))

}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	cursos := Cursos{
		{"Go", 40},
		{"Java", 400},
		{"Python", 50},
	}
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(w, cursos)
	if err != nil {
		panic(err)
	}
}
