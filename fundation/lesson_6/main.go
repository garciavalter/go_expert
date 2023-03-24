package main

import "fmt"

func main() {
	// declarando um slice e inicializando com valores, poderia ser inicializado como vazio {}.
	s := []int{2, 4, 6, 8, 10}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// colocando dois pontos antes da posiçao (len) faz com que o que estiver a direita seja dropado do slice.
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	//dropando a partir da posição 2.
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
	//dropando a partir da posição 4.
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
	// colocando dois pontos depois da posiçao (len) faz com que o
	// que estiver a esquerda seja dropado do slice, nesse caso a capacidade é diminuida.
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])
	s = append(s, 12)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
