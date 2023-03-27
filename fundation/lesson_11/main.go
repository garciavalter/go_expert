package main

import "fmt"

// declarando uma struct
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", wesley.Nome, wesley.Idade, wesley.Ativo)
	// alterando um parametro dentro de uma struct
	wesley.Ativo = false
	fmt.Println(wesley)
}
