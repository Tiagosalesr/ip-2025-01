package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)
	if n == 0 {
		fmt.Print("O número é nulo.")
	} else if n > 0 {
		fmt.Print("O número é positivo.")
	} else {
		fmt.Print("O número é negativo.")
	}
}
