package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três"}
	//no ranged o primeiro valor é o índice, o segundo os valores.
	for k, v := range numeros {
		println(k, v)
	}

	//usando o for como estrutura do while.
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	//loop infinito
	for {
		println("Dormammu! I've come to bargain!")
	}
}
