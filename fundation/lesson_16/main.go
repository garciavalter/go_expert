package main

func somaPonteiro(a, b *int) int {
	//alterando o valor armazenado na memória pela função
	*a = 50
	return *a + *b
}

func somaValores(a, b int) int {
	// altera somente no escopo local
	a = 50
	return a + b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	// passando cópia dos valores
	somaValores(minhaVar1, minhaVar2)
	//não altera valor de minhaVar1
	println(minhaVar1)
	//passando o endereço na memória
	somaPonteiro(&minhaVar1, &minhaVar2)
	// altera valor de minhaVar1
	println(minhaVar1)
}
