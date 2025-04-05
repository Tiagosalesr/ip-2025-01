package main

import "fmt"

func main() {
	var nota float64
	var conceito string

	fmt.Print("Digite a nota do aluno: ")
	fmt.Scan(&nota)
	if nota <= 10 && nota >= 9 {
		conceito = "A"
	} else if nota < 9 && nota >= 7.5 {
		conceito = "B"
	} else if nota < 7.5 && nota >= 6 {
		conceito = "C"
	} else if nota < 6 && nota >= 0 {
		conceito = "D"
	} else {
		conceito = "Nota inválida"
	}
	fmt.Printf("NOTA = %.2f CONCEITO = %s", nota, conceito)
}
