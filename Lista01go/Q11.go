package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite um número: ")
	fmt.Scan(&x)
	if x%3 > -1 && x%3 < 1 && x%5 > -1 && x%5 < 1 {
		fmt.Println("O NUMERO E DIVISIVEL")
	} else {
		fmt.Println("O NUMERO NAO E DIVISIVEL")
	}
}
