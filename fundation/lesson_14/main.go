package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// Cria a interface e implementa automaticamente como Pessoa para todas as structs que tenham o método desativar
type Pessoa interface {
	//A interface permite passar apenas métodos.
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado, ativo: %t", c.Nome, c.Ativo)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()

}

type Empresa struct {
	Nome string
}

// Qualquer struct que tiver o método desativar implementar a interface automaticamente.
func (e Empresa) Desativar() {

}
func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	minhaEmrpesa := Empresa{
		Nome: "Minha Empresa",
	}

	wesley.Desativar()
	Desativacao(wesley)
	Desativacao(minhaEmrpesa)
}
