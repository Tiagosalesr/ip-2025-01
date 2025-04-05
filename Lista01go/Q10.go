package main

import "fmt"

func main() {
	var a, b, c, d, det float64
	fmt.Print("Defina a: ")
	fmt.Scan(&a)
	fmt.Print("Defina b: ")
	fmt.Scan(&b)
	fmt.Print("Defina c: ")
	fmt.Scan(&c)
	fmt.Print("Defina d: ")
	fmt.Scan(&d)
	det = (a * d) - (b * c)
	fmt.Printf("O VALOR DO DETERMINANTE É = %.2f", det)
}
