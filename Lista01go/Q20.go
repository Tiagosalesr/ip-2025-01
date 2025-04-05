package main

import "fmt"

func main() {
	var h, min, seg int
	fmt.Print("Digite a hora: ")
	fmt.Scan(&h)
	fmt.Print("Digite os minutos: ")
	fmt.Scan(&min)
	fmt.Print("Digite os segundos: ")
	fmt.Scan(&seg)
	h = h * 3600
	min = min * 60
	fmt.Printf("O TEMPO EM SEGUNDOS É: %d\n", h+min+seg)
}
