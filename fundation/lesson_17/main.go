package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo int
}

// Criando structs com endereçamento na memória
func NewConta() *Conta {
	return &Conta{saldo: 0}
}

// Manipula a struct como uma cópia, não alterando a original.
func (c Cliente) andou() {
	c.nome = "Wesley Williams"
	fmt.Printf("O cliente %v andou", c.nome)
}

// Alternado o valor da struct na memória passando o ponteiro como argumento no método
func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {

	wesley := Cliente{
		nome: "Wesley",
	}
	wesley.andou()

	//nome não alterado
	fmt.Printf("O valor da struct com nome %v", wesley.nome)

	conta := Conta{
		saldo: 100,
	}
	conta.simular(200)
	// saldo alterado
	println(conta.saldo)
}
