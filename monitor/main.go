package main

import (
	"fmt"
	"os"
)

func main() {
	exibeMenu()
	option := captaOpcao()

	switch option {
	case 1:
		fmt.Println("Monitoramento Iniciado")
	case 2:
		fmt.Println("Logs exibidos")
	case 0:
		os.Exit(0)
	default:
		fmt.Println(option, "Não é esperado")
		os.Exit(-1)
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func captaOpcao() int {
	var option int
	fmt.Scan(&option)
	fmt.Println("Foi escolhida a opção", option, "no endereco de memoria", &option)

	return option
}
