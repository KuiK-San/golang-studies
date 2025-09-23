package contas

type ContaBanco struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaBanco) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorSaque
		return "Valor Sacado"
	}

	return "Saldo Insuficiente"
}

func (c *ContaBanco) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.Saldo += valorDeposito
		return "Valor Depositado", c.Saldo
	}
	return "Valor invalido", c.Saldo
}

func (c *ContaBanco) Transferir(valorTransferencia float64, contaDestino *ContaBanco) bool {
	if valorTransferencia <= 0 {
		return false
	}
	if valorTransferencia < c.Saldo {
		c.Sacar(valorTransferencia)
		contaDestino.Depositar(valorTransferencia)
		return true
	}

	return false
}
