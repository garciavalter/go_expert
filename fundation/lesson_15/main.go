package main

// trabalhando com ponteiros
func main() {
	// Memória  -> Endereço -> Valor = 10
	a := 10
	// E comercial antes da variável acessa o endereço da memória que armazena o valor.
	println(&a)
	// Toda vez que utilizamos asterisco estamos apontando para um endereço na memória.
	var ponteiro *int = &a
	// A variávle aponta para um ponteiro que tem um endereço na memória que tem um valor.
	println(ponteiro)
	// Quando alteramos o ponteiro alteramos o valor no endereço da memória
	*ponteiro = 20
	println(a) //vai printar 20
	//declarando ponteiros via shorthand
	b := &a
	println(b)
	//acessando o valor armazenado no endereço da memória (deep reference)
	*b = 30
	println(*b)
	println(a)

}
