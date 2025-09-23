package main

import "fmt"

type ContaBanco struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaBanco) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Valor Sacado"
	}

	return "Saldo Insuficiente"
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

func main() {
	conta1 := ContaBanco{titular: "Guilherme", saldo: 500}
	conta2 := ContaBanco{titular: "Guilherme2", saldo: 500}

	status := conta2.Transferir(700, &conta1)

	fmt.Println(status)
	fmt.Println(conta1)
	fmt.Println(conta2)
}
