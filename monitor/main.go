package main

import "fmt"

func main() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")

	var option int
	fmt.Scan(&option)
	fmt.Println("Foi escolhida a opção", option, "no endereco de memoria", &option)
}
