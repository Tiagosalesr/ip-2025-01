package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("Digite dois números inteiros:")
	fmt.Scanln(&x, &y)

	if x%2 == 0 {
		for i := 0; i < y; i++ {
			fmt.Print(x+(i*2), " ")
		}
	} else {
		fmt.Println("PRIMEIRO NÚMERO NÃO É PAR")
	}
}
