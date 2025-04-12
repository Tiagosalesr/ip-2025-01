package main

import "fmt"

func main() {
	var (
		idade  int
		classe string
	)
	for true {
		fmt.Print("Digite sua idade: ")
		fmt.Scan(&idade)
		if idade < 0 {
			fmt.Println("Idade inválida")
		} else {
			break
		}
	}
	if idade < 16 {
		classe = "Não-eleitor"
	} else if idade >= 16 && idade < 18 || idade > 65 {
		classe = "Eleitor facultativo"
	} else {
		classe = "Eleitor obrigatório"
	}
	fmt.Print("A sua classe eleitoral é: ", classe)
}
