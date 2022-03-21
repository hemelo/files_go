package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"strconv"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()
	for { 
		exibeMenu()

		comando := leComando()
		fmt.Println("")

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando...")
			os.Exit(-1)
		}
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func exibeIntroducao() {
	nome := "Miyazaki"
	versao := 1.1
	fmt.Println("Oi", nome)
	fmt.Println("Este programa está na versão", versao)
}

func iniciarMonitoramento(){
	fmt.Println("Monitorando...")

	sites := leSitesDoArquivo()

	for i := 1; i <= monitoramentos; i++ {
		fmt.Println("Roda de testes de número", i)

		for _, site := range sites{
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}	
		
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "apresenta problemas.", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	arquivo, err := os.Open("sites.txt")

	if err != nil { 
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(arquivo)

	var sites []string 

	for { 
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF { 
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_APPEND | os.O_RDWR | os.O_CREATE, 0666)
	
	if err != nil { 
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + "- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	
	if err != nil { 
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}