package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Escreva os dois números separados por um espaço: ")
	fmt.Scan(&n1, &n2)
	if n1%n2 == 0 {
		fmt.Print("O primeiro número é divisível pelo segundo.")
	} else {
		fmt.Print("O primeiro número não é divisível pelo segundo.")
	}
}
