package main

import (
	"fmt"
	"net/http"
	"os"
)


func main(){
	hello()
	exibirMenu()
	comando := lerComando()

	switch comando{
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não sei o que tu tá dizendo")
		os.Exit(-1)
	}
}

func hello(){
	fmt.Println("Qual seu nome?")
	var nome string
	fmt.Scan(&nome)
	fmt.Println("Olá,",nome)
}

func exibirMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")	
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func iniciarMonitoramento(){
	fmt.Println("Iniciando monitoramento")
	site := "https://ilovemyjob.com.br/"
	resp, _ := http.Get(site)
	if resp.StatusCode == 200{
		fmt.Println("O site", site, "está no ar.")
	}else{
		fmt.Println("Fudeu")
	}
}