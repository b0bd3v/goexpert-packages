package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"` // Tag para alterar o nome do campo no JSON
	Saldo  int `json:"s"`
	//	Se passar o - ele ignora o campo
}

func main() {
	conta := Conta{Numero: 123, Saldo: 1000}
	res, err := json.Marshal(conta)

	if err != nil {
		panic(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)

	if err != nil {
		panic(err)
	}

	jsonPuro := `{"n":123,"s":1000}`

	var contaX Conta

	// Unica forma de alterar o valr da variavel e passando o endereco de memoria
	err = json.Unmarshal([]byte(jsonPuro), &contaX)

	if err != nil {
		panic(err)
	}

	println(contaX.Saldo)
}
