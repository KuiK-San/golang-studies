package contas

import "github.com/kuik/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Valor Sacado"
	}

	return "saldo Insuficiente"
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Valor Depositado", c.saldo
	}
	return "Valor invalido", c.saldo
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}
