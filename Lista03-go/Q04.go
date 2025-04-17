package main

import (
	"fmt"
)

func main() {
	var num int

	for {
		fmt.Print("Digite um número inteiro (0 ou negativo para sair): ")
		fmt.Scanln(&num)

		if num <= 0 {
			fmt.Println("Programa encerrado.")
			break
		}

		if ehQuadradoPerfeito(num) {
			fmt.Printf("%d é um quadrado perfeito.\n", num)
		} else {
			fmt.Printf("%d não é um quadrado perfeito.\n", num)
		}
	}
}

func ehQuadradoPerfeito(n int) bool {
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return true
		}
	}
	return false
}
