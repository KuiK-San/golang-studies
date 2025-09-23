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

func main() {
	conta1 := ContaBanco{titular: "Guilherme", saldo: 500}

	fmt.Println(conta1.saldo)

	fmt.Println(conta1.Sacar(0))
	fmt.Println(conta1.saldo)
}
