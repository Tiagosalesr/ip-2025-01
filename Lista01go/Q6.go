package main

import "fmt"

func main() {
	var f, n1 float64
	var array []float64

	fmt.Print("Digite o número de temperaturas a converter: ")
	fmt.Scan(&n1)
	for i := 0; i < int(n1); i++ {
		fmt.Print("Digite a temperatura em Fahrenheit: ")
		fmt.Scan(&f)
		array = append(array, (f-32)*5/9)
	}
	for i := 0; i < int(n1); i++ {
		far := ((array[i]) * (9.0 / 5.0)) + 32
		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A  %.2f CELSIUS\n", far, array[i])
	}
}
