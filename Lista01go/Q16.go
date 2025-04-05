package main

import "fmt"

func main() {
	var salario, reajuste float64
	fmt.Println("Digite o salário atual:")
	fmt.Scan(&salario)
	if salario <= 300 {
		reajuste = salario * 1.5
	} else {
		reajuste = salario * 1.3
	}
	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", reajuste)
}
