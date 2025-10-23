package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

// criando func  para sacar
func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(ValorDoDeposito float64) (string, float64) {
	if ValorDoDeposito > 0 {
		c.saldo += ValorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito inválido", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
