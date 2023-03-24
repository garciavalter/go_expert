package main

import (
	"errors"
	"fmt"
)

// tratando erros em go
func main() {
	valor, err := sum2(50, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

// Declarando função com retorno
func sum(a int, b int) int {
	return a + b
}

// Declarando função com parametros de entrada de mesmo tipo
func sum1(a, b int) int {
	return a + b
}

// Declarando função com retorno de mais de um valor
func sum2(a, b int) (int, error) {
	if a+b >= 50 {
		// No go não há try catch, então é necessário retornar um erro
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
