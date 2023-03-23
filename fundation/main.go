package main

// constante de escopo global
const a = "Hello, World!"

// variáveis de escopo global
// tudo no go é fortemente tipado, ou seja, não é possível atribuir um valor de um tipo diferente ao declarado
var (
	b bool = true //variável declarada com valor inicial
	c int         //variável declarada sem valor inicial
	d string
	e float64
)

func main() {
	//variável de escopo local
	a := "string" //variáveis locais declaras e não utilizadas geram erro.
	// no go podemos usar := para declarar e atribuir um valor a uma variável com inferência de tipo automática.
	b = true //variáveis globais podem ser alteradas dentro de funções
	println(a)
}
