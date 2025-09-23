package main

import (
	"fmt"

	"github.com/kuik/clientes"
	"github.com/kuik/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	cliente1 := clientes.Titular{
		Nome:      "Guilherme",
		CPF:       "123.456.789-10",
		Profissao: "DevOps",
	}
	conta1 := contas.ContaBanco{
		Titular: cliente1,
	}
	conta1.Depositar(500)

	cliente2 := clientes.Titular{
		Nome:      "Guilherme2",
		CPF:       "123.456.789-10",
		Profissao: "Desenvolvedor",
	}
	conta2 := contas.ContaPoupanca{
		Titular: cliente2,
	}
	conta2.Depositar(500)

	PagarBoleto(&conta1, 10)
	PagarBoleto(&conta2, 20)

	fmt.Println(conta1.GetSaldo())
	fmt.Println(conta2.GetSaldo())
}
