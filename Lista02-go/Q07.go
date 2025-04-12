package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c, menor, inter, maior int
	fmt.Print("Escreva  três números distintos separados por um espaço: ")
	fmt.Scan(&a, &b, &c)
	arr := []int{a, b, c}
	sort.Ints(arr)
	menor = arr[0]
	inter = arr[1]
	maior = arr[2]
	fmt.Print(menor, inter, maior)
}
