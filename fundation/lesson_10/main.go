package main

import (
	"fmt"
)

func main() {
	//declarando funções anônimas no go
	total := func() int {
		return sum(1, 2, 3) * 2
	}()
	valor := sum(10, 20, 30)
	fmt.Printf("O valor é %d\n", valor)
	fmt.Printf("O total é %d\n", total)

}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
