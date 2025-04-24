package main

import "fmt"

func main() {
	var (
		num     int
		array   = []int{}
		pares   = []int{}
		impares = []int{}
	)
	for i := 1; i <= 10; i++ {
		fmt.Printf("Digite o %dº termo: ", i)
		fmt.Scan(&num)
		array = append(array, num)
	}
	somaPares := 0
	qntdImpares := 0
	for _, v := range array {
		if v%2 == 0 {
			pares = append(pares, v)
			somaPares += v
		}
		if v%2 != 0 {
			impares = append(impares, v)
			qntdImpares += 1
		}
	}
	fmt.Printf("Os números pares são: %d\n", pares)
	fmt.Printf("A soma dos números pares é: %d\n", somaPares)
	fmt.Printf("Os números ímpares são: %d\n", impares)
	fmt.Printf("A quantidade de números ímpares é: %d\n", qntdImpares)
}
