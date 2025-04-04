package main

import "fmt"

func main() {
	var (
		conta, consumo, valor float64
		tipo                  string
	)
	fmt.Println("Digite o número da conta, o consumo em metros cúbicos e o tipo da conta:")
	fmt.Scanln(&conta, &consumo, &tipo)
	valor = 0
	if tipo == "R" {
		valor += 5 + (consumo * 0.05)
	} else if tipo == "I" && consumo <= 100 {
		valor += 800
	} else if tipo == "I" && consumo > 100 {
		valor += 800 + ((consumo - 100) * 0.04)
	} else if tipo == "C" && consumo <= 80 {
		valor += 500
	} else if tipo == "C" && consumo > 80 {
		valor += 500 + ((consumo - 80) * 0.03)
	} else if tipo != "C" && tipo != "I" && tipo != "R" {
		fmt.Println("Tipo de conta inválido.")
	}
	fmt.Printf("CONTA: %.0f\nVALOR DA CONTA: R$ %.2f\n", conta, valor)
}
