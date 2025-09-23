package main

import (
	"fmt"

	"github.com/kuik/contas"
)

func main() {
	conta1 := contas.ContaBanco{Titular: "Guilherme", Saldo: 500}
	conta2 := contas.ContaBanco{Titular: "Guilherme2", Saldo: 500}

	status := conta2.Transferir(700, &conta1)

	fmt.Println(status)
	fmt.Println(conta1)
	fmt.Println(conta2)
}
