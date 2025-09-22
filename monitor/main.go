package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	for {
		exibeMenu()
		option := captaOpcao()

		switch option {
		case 1:
			monitoraSite()
		case 2:
			fmt.Println("Logs exibidos")
		case 0:
			os.Exit(0)
		default:
			fmt.Println(option, "Não é esperado")
			os.Exit(-1)
		}
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

func monitoraSite() {
	fmt.Println("Monitorando...")
	site := "https://httpbin.org/status/200"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: tr,
	}

	resp, _ := client.Get(site) // http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "funcionando normalmente.")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
