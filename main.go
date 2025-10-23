package main

import (
	"banco/contas"
	"fmt"
)

// utilizando no struct no main
func main() {

	contaDaSilva := contas.ContaCorrente{}
	contaDaSilva.Titular = "Silva"
	contaDaSilva.Saldo = 500

	fmt.Println(contaDaSilva.Saldo)

	fmt.Println(contaDaSilva.Sacar(-1000))

	fmt.Println(contaDaSilva.Saldo)

	fmt.Println(contaDaSilva.Depositar(1000))

	status, valor := contaDaSilva.Depositar(200)
	fmt.Println(status, valor)

	contaDoNeto := contas.ContaCorrente{Titular: "Neto", Saldo: 400}
	contaDoMarcos := contas.ContaCorrente{Titular: "Marcos", Saldo: 400}

	fmt.Println(contaDoMarcos)
	fmt.Println(contaDoNeto)

	flag := contaDoMarcos.Transferir(-200, &contaDoNeto)

	fmt.Println(flag)
	fmt.Println(contaDoNeto.Saldo)
}
