package main

import (
	"fmt"
)

// criando uma estrutura em go
type ContaCorrente struct {
	titular     string
	agencia     int
	numeroConta int
	saldo       float64
}

// utilizando no struct no main
func main() {

	conta := ContaCorrente{titular: "Goku", agencia: 111,
		numeroConta: 123456, saldo: 142.50}
	conta2 := ContaCorrente{titular: "Goku", agencia: 111,
		numeroConta: 123456, saldo: 142.50}
	//como se fosse o equals, pois compara o conteúdo.
	fmt.Println(conta == conta2)
	// conta2 := ContaCorrente{"Neto", 111, 123456, 150.00}
	// fmt.Println("Essa é outra conta: ", conta2)

	// //exemplo de como utilizar ponteiros
	// var contaDoNeto *ContaCorrente
	// contaDoNeto = new(ContaCorrente)
	// contaDoNeto.titular = "Neto"

	// fmt.Println(*contaDoNeto)

}
