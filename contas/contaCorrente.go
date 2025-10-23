package contas

// criando uma estrutura em go
type ContaCorrente struct {
	Titular     string
	Agencia     int
	NumeroConta int
	Saldo       float64
}

// criando func  para sacar
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(ValorDoDeposito float64) (string, float64) {
	if ValorDoDeposito > 0 {
		c.Saldo += ValorDoDeposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "Valor do depósito inválido", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(ValorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if ValorDaTransferencia < c.Saldo && ValorDaTransferencia > 0 {
		c.Saldo -= ValorDaTransferencia
		contaDestino.Depositar(ValorDaTransferencia)
		return true
	} else {
		return false
	}
}
