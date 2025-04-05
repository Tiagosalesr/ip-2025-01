package main

import (
	"fmt"
	"math"
)

func main() {
	var a, h, v, ab, raiz float64
	fmt.Println("DIGITE A ALTURA:")
	fmt.Scan(&h)
	fmt.Println("DIGITE O COMPRIMENTO  DA ARESTA:")
	fmt.Scan(&a)
	raiz = math.Sqrt(3)
	ab = (3 * a * a * raiz) / (2)
	v = (ab * h) / 3
	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CÚBICOS\n", v)
}
