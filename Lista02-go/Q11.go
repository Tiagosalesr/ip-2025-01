package main

import "fmt"

func main() {
	var x, y float64
	fmt.Print("Defina x: ")
	fmt.Scan(&x)
	y = 8 / (2 - x)
	fmt.Print("f(x) = ", y)
}
