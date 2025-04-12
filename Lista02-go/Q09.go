package main

import "fmt"

func main() {
	var valorCompra, valorVenda, lucro float64
	fmt.Println("Escreva o valor da compra:")
	fmt.Scanln(&valorCompra)
	if valorCompra < 10 && valorCompra > 0 {
		lucro = 1.7
	} else if valorCompra >= 10 && valorCompra < 30 {
		lucro = 1.5
	} else if valorCompra >= 30 && valorCompra < 50 {
		lucro = 1.4
	} else if valorCompra >= 50 {
		lucro = 1.3
	}
	valorVenda = valorCompra * lucro
	if valorVenda > 0 {
		fmt.Printf("O valor da venda é: R$%.2f", valorVenda)
	} else {
		fmt.Print("Valor de compra inválido")
	}
}
