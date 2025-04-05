package main

import "fmt"

func main() {
	var a1, razao, numeroElementos, soma int
	soma = 0
	fmt.Println("Digite oprimeiro termo da P.A., a razão e o número de elementos:")
	fmt.Scan(&a1, &razao, &numeroElementos)
	var mapa = make([]int, numeroElementos)
	for i := 0; i < numeroElementos; i++ {
		mapa[i] = a1 + (i * razao)

	}
	for i := 0; i < numeroElementos; i++ {
		soma += mapa[i]
	}
	fmt.Printf("Soma dos itens da P.A. = %d ", soma)

}
