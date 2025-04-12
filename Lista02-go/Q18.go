package main

import "fmt"

func main() {
	var (
		dia       int
		categoria string
		preco     float64
	)
	for true {
		fmt.Print("Digite o dia da semana, a categoria do DVD e seu preço original: ")
		fmt.Scan(&dia, &categoria, &preco)
		if dia > 7 || dia < 1 || categoria != "comum" && categoria != "lançamento" || preco < 1 {
			fmt.Println("Dados inválidos")
		} else {
			break
		}
	}
	if dia == 2 && categoria == "comum" || dia == 3 && categoria == "comum" || dia == 5 && categoria == "comum" {
		preco *= 0.6
	} else if dia == 2 && categoria == "lançamento" || dia == 3 && categoria == "lançamento" || dia == 5 && categoria == "lançamento" {
		preco *= 1.15
		preco *= 0.6
	} else if dia == 4 && categoria == "lançamento" || dia == 6 && categoria == "lançamento" || dia == 7 && categoria == "lançamento" || dia == 1 && categoria == "lançamento" {
		preco *= 1.15
	} else {
		preco = preco
	}
	fmt.Printf("O preço do DVD é de: R$ %.2f", preco)

}
