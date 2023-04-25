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
	curso := Curso{"Go", 40}
	//trantando os erros de uma vez só

	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
