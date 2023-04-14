package main

import (
	"encoding/json"
	"os"
)

// Uso de tags para json
type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{
		Numero: 1,
		Saldo:  100,
	}
	//Usando o Marshal o valor é armazenado na memória
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))
	//usando enconder é feito um stream direto para o destino
	//pode ser a response do webserver, gravar arquivos
	enconder := json.NewEncoder(os.Stdout)
	enconder.Encode(conta)

	jsonPuro := []byte(`{"Numero":2,"Saldo":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)
	var contaY Conta
	jsonIncorreto := []byte(`{"n":2,"s":200}`)
	err = json.Unmarshal(jsonIncorreto, &contaY)
	if err != nil {
		println(err)
	}
	println(contaY.Saldo)

}
