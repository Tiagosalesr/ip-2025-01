package main

import "fmt"

func main() {
	var cel, fahr, pol, mm float64
	fmt.Println("Digite a temperatura em fahrenheit: ")
	fmt.Scanln(&fahr)
	cel = (fahr - 32) * 5 / 9
	fmt.Println("Digite o volume de chuva em polegadas")
	fmt.Scanln(&pol)
	mm = pol * 25.4
	fmt.Printf("O VALOR EM CELSIUS = %.2f\n A QUANTIDADE DE CHUVA É %.2f mm\n", cel, mm)
}
