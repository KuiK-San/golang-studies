package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	for {
		exibeMenu()
		option := captaOpcao()

		switch option {
		case 1:
			monitoraSites()
		case 2:
			imprimeLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println(option, "Não é esperado")
			os.Exit(-1)
		}
		fmt.Println("")
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

func monitoraSites() {
	fmt.Println("Monitorando...")
	sites := getSitesFromFile()

	for i := 0; i <= monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	/* tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: tr,
	} */
	resp, err := http.Get(site) // client.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro na requisição:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "funcionando normalmente.")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func getSitesFromFile() []string {
	var sites []string
	arquivo, err := os.Open("./sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(arquivo)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Ocorreu um erro", err)
		}

	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao executar:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := os.ReadFile("./log.txt")

	if err != nil {
		fmt.Println("Deu erro: ", err)
	}

	fmt.Println(string(arquivo))

}
