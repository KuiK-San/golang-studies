package main

import "fmt"

type ContaBanco struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	conta1 := ContaBanco{titular: "Guilherme", numeroAgencia: 333, numeroConta: 123123, saldo: 125.5}
	conta2 := ContaBanco{"Teste2", 222, 123333, 123.5}

	fmt.Println(conta1)
	fmt.Println(conta2)

	var conta3 *ContaBanco
	conta3 = new(ContaBanco)
	conta3.titular = "Teste3"

	fmt.Println(*conta3)
}
