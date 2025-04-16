package main

import "fmt"

func main (){
	var list[] int
	soma := 0
	for i := 50; i < 71; i += 2{
		list = append(list, i)
		soma += i 
	}
	media := soma / len(list)
	fmt.Printf("A soma dos números pares de 50 a 70 é: %d e a média aritmètica desses números é: %d\n", soma, media)
}
