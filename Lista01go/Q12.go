package main

import "fmt"

func main() {
	var horas, valor, i int

	fmt.Println("Digite o número de horas:")
	fmt.Scan(&horas)
	if horas%3 == 0 {
		valor = ((horas / 3) * 10)
	} else {
		for horas%3 != 0 {
			horas--
			i++
		}
		valor = ((horas / 3) * 10) + (i * 5)
	}
	fmt.Printf("O valor a ser pago é: R$ %.2f\n", float64(valor))
}
