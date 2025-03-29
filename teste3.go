package main

import "fmt"

func main () {

	var soma int

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		soma = soma + i
	}
	fmt.Println(soma)
}