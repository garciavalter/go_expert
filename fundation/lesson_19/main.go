package main

import "fmt"

func main() {
	var minhaVar interface{} = "Wesley Williams"
	//usando println é necessário informar o tipo para imprimir
	println(minhaVar.(string))
	// fazendo asserções de tipos
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
	//não passar o "ok" causa um panic em caso de falha
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)
}
