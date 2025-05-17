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

func main() {
	conta := ContaCorrente{titular: "Goku", agencia: 111,
		numeroConta: 123456, saldo: 142.50}
	fmt.Println(conta)
	conta2 := ContaCorrente{"Neto", 111, 123456, 150.00}
	fmt.Println("Essa é outra conta: ", conta2)
}
