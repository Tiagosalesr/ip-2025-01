package main

import "fmt"

func main() {
	var at, al, ac, pi, r, a, valor float64
	fmt.Print("Digite o raio: ")
	fmt.Scan(&r)
	fmt.Print("Digite a altura: ")
	fmt.Scan(&a)
	pi = 3.14159
	al = a * 2 * pi * r
	ac = pi * r * r
	at = al + (2 * ac)
	valor = at * 100
	fmt.Printf("O VALOR DO CUSTO E = %.2f", valor)
}
