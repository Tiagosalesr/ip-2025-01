package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)
	if n < 0 {
		fmt.Print(math.Pow(n, 2))
	} else {
		fmt.Print(math.Sqrt(n))
	}
}
