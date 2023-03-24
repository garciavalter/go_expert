package main

import (
	"fmt"
)

func main() {
	valor := sum(10, 20, 30)
	fmt.Println(valor)
}

// Declarando função variática
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
