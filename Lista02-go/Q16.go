package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, delta, raiz1, raiz2 float64
	fmt.Print("Digite a, b e c da equação (separados por um espaço): ")
	fmt.Scan(&a, &b, &c)
	delta = (b * b) - (4 * a * c)
	if a == 0 {
		fmt.Println("A equação não é do segundo grau. O coeficiente A não pode ser zero.")
	} else if delta < 0 {
		fmt.Print("RAÍZES IMAGINÁRIAS")
	} else if delta == 0 {
		raiz1 = (-b) / (2 * a)
		fmt.Print("RAIZ ÚNICA:", raiz1)
	} else {
		raiz1 = ((-b) + math.Sqrt(delta)) / (2 * a)
		raiz2 = ((-b) - math.Sqrt(delta)) / (2 * a)
		fmt.Printf("RAÍZES DISTINTAS: %.2f e %.2f", raiz1, raiz2)
	}
}
