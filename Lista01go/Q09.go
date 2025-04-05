package main

import "fmt"

func main() {
	var delta, a, b, c float64
	fmt.Print("Digite a: ")
	fmt.Scan(&a)
	fmt.Print("Digite b: ")
	fmt.Scan(&b)
	fmt.Print("Digite c: ")
	fmt.Scan(&c)
	delta = (b * b) - (4 * a * c)
	fmt.Printf("O VALOR DE DELTA E = %.2f", delta)
}
