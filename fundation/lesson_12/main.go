package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// Compondo structs
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	//acessando informações encadeadas
	wesley.Endereco.Cidade = "São Paulo"
	//acessando por shorthand
	wesley.Cidade = "Rio de Janeiro"
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", wesley.Nome, wesley.Idade, wesley.Ativo)

	// alterando um parametro dentro de uma struct
	wesley.Ativo = false
	fmt.Println(wesley)
}
