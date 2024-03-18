package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s" validate:"gt=0"`
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

	jsonPuro := []byte(`{"n":2, "s": 200}`)
	var contaSecundaria Conta
	json.Unmarshal(jsonPuro, &contaSecundaria)
	if err != nil {
		println(err)
	}
	println(contaSecundaria.Saldo)
}
