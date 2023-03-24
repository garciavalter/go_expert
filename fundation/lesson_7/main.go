package main

import "fmt"

func main() {
	//iniciando um map com dados
	salarios := map[string]int{"João": 1500, "Maria": 3000, "Pedro": 2000}
	//deletando uma chave
	delete(salarios, "João")
	//adicionando uma chave
	salarios["Tiago"] = 5000
	//iniciando um map vazio com o make
	sal := make(map[string]int)
	sal["João"] = 1500
	sal["Maria"] = 3000
	sal["Pedro"] = 2000
	//percorrendo um map
	for nome, salario := range salarios {
		fmt.Printf("Nome: %s Salário: %d\n", nome, salario)
	}
	//usando blank identifier para ignorar um parâmetro
	for _, salario := range salarios {
		fmt.Printf("Salário: %d\n", salario)
	}
}
