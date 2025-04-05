package main

import "fmt"

func main() {
	var (
		soma, a float64
		n       int
	)
	fmt.Print("Digite um número: ")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		a = 1 / float64(i)
		soma += a
	}
	fmt.Printf("A soma dos inversos de 1 a %d é: %.6f\n", n, soma)
}
