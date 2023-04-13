package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}
	f.Close()
	fmt.Printf("Arquivo criado com sucesso, o tamanho do arquivo é %d bytes\n", tamanho)
	//ReadFile não abre o arquivo, faz a leitura direto, por isso não é necessário close.
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
