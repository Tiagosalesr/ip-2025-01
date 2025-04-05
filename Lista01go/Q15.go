package main

import "fmt"

func main() {

	var n, n1 int
	var results []string // Slice para armazenar os resultados

	fmt.Println("Digite um número:")
	fmt.Scanln(&n)
	if 5 < n && n < 2000 {

		for n > 0 {
			if n%2 == 0 {
				n1 = n * n
				results = append(results, fmt.Sprintf("%d^2 = %d", n, n1)) // Armazena o resultado
			}
			n = n - 2
		}

		// Exibe os resultados em ordem crescente
		for i := len(results) - 1; i >= 0; i-- {
			fmt.Println(results[i])
		}
	}
}
