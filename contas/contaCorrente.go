package contas

import "banco/clientes"

// criando uma estrutura em go
type ContaCorrente struct {
	Titular     clientes.Titular
	Agencia     int
	NumeroConta int
	saldo       float64
}

// criando func  para sacar
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(ValorDoDeposito float64) (string, float64) {
	if ValorDoDeposito > 0 {
		c.saldo += ValorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito inválido", c.saldo
	}
}

func (c *ContaCorrente) Transferir(ValorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if ValorDaTransferencia < c.saldo && ValorDaTransferencia > 0 {
		c.saldo -= ValorDaTransferencia
		contaDestino.Depositar(ValorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
