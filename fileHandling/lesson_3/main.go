package main

import (
	"fmt"
)

func main() {
	//O defer move a execução para o final do arquivo
	//útil para fechar stream de dados
	defer fmt.Println("Primeira linha")
	fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")
}
