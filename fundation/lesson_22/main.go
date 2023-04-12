package main

import (
	"fmt"
	//Para usar módulos próprios é necessário rodar go mod init passando um nome de pacote.
	//isso criará o arquivo go.mod responsável por gerenciar pacotes
	"go/expert/lesson21/matematica"
)

func main() {
	soma := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro.Andar())
	fmt.Println("O resultado é :", soma)
	fmt.Println(matematica.A)

}
