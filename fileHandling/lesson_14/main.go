package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	cursos := Cursos{
		{"Go", 40},
		{"Java", 400},
		{"Python", 50},
	}
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, cursos)
	if err != nil {
		panic(err)
	}

}
