package main

import (
	"fmt"
	//Para usar módulos próprios é necessário rodar go mod init passando um nome de pacote.
	//isso criará o arquivo go.mod responsável por gerenciar pacotes
	"go/expert/lesson21/matematica"
	//go mod tidy é capaz de baixar pacotes e instalar automaticamente
	// esse comando também limpa o go.mod e go.sum para apenas pacotes em uso
	"github.com/google/uuid"
)

func main() {
	soma := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro.Andar())
	fmt.Println("O resultado é :", soma)
	fmt.Println(matematica.A)
	fmt.Println(uuid.New())
}
