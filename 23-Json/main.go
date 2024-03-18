package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int
	Saldo  int
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"Numero":2, "Saldo": 200}`)
	var contaSecundaria Conta
	json.Unmarshal(jsonPuro, &contaSecundaria)
	if err != nil {
		println(err)
	}
	println(contaSecundaria.Saldo)
}
