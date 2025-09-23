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

func main() {
	conta1 := ContaBanco{titular: "Guilherme", saldo: 500}

	fmt.Println(conta1.saldo)

	fmt.Println(conta1.Depositar(-500))
	fmt.Println(conta1.saldo)
}
