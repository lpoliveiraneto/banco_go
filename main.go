package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

// utilizando no struct no main
func main() {
	//exemplos sem a composição do struct
	// contaDaSilva := contas.ContaCorrente{}
	// contaDaSilva.Titular = "Silva"
	// contaDaSilva.Saldo = 500

	// fmt.Println(contaDaSilva.Saldo)

	// fmt.Println(contaDaSilva.Sacar(-1000))

	// fmt.Println(contaDaSilva.Saldo)

	// fmt.Println(contaDaSilva.Depositar(1000))

	// status, valor := contaDaSilva.Depositar(200)
	// fmt.Println(status, valor)

	// contaDoNeto := contas.ContaCorrente{Titular: "Neto", Saldo: 400}
	// contaDoMarcos := contas.ContaCorrente{Titular: "Marcos", Saldo: 400}

	// fmt.Println(contaDoMarcos)
	// fmt.Println(contaDoNeto)

	// flag := contaDoMarcos.Transferir(200, &contaDoNeto)

	// fmt.Println(flag)
	// fmt.Println(contaDoNeto.Saldo)

	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Bruno",
		CPF:       "033.787.043-81",
		Profissao: "puta"},
		NumeroConta: 12346}

	fmt.Println(contaDoBruno.Depositar(1000))

	PagarBoleto(&contaDoBruno, 500)

	fmt.Println(contaDoBruno.ObterSaldo())

	contaMinha := contas.ContaPoupanca{}
	contaMinha.Depositar(200)

	fmt.Println(contaMinha)
	PagarBoleto(&contaMinha, 100)
	fmt.Println(contaMinha)
}
