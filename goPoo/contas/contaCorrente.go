package contas

import "github.com/kuik/clientes"

type ContaBanco struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaBanco) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Valor Sacado"
	}

	return "saldo Insuficiente"
}

func (c *ContaBanco) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Valor Depositado", c.saldo
	}
	return "Valor invalido", c.saldo
}

func (c *ContaBanco) Transferir(valorTransferencia float64, contaDestino *ContaBanco) bool {
	if valorTransferencia <= 0 {
		return false
	}
	if valorTransferencia < c.saldo {
		c.Sacar(valorTransferencia)
		contaDestino.Depositar(valorTransferencia)
		return true
	}

	return false
}

func (c *ContaBanco) GetSaldo() float64 {
	return c.saldo
}
