package main

// constante de escopo global
const a = "Hello, World!"

type ID int //criando um tipo ID que é um int

var (
	b bool = true
	c int
	d string
	e float64
	f ID = 1
)

func main() {

	a := "string"

	b = true
	println(a)
}
