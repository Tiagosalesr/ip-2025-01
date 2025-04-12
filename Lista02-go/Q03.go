package main

import "fmt"

func main() {
	var n1, n2, soma int
	fmt.Print("Escreva os dois números separados por um espaço: ")
	fmt.Scan(&n1, &n2)
	soma = n1 + n2
	if soma > 20 {
		fmt.Print(soma + 8)
	} else {
		fmt.Print(soma - 5)
	}
}
