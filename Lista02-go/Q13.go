package main

import "fmt"

func main() {
	var n int
	for true {
		fmt.Print("Digite o número de 3 dígitos: ")
		fmt.Scan(&n)
		if n < 100 || n > 999 {
			fmt.Println("Número inválido")
		} else {
			break
		}
	}
	dezena := (n / 10) % 10
	fmt.Println("O algorismo da casa das dezenas é:", dezena)
}
