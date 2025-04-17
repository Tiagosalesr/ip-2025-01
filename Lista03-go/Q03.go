package main

import f "fmt"

func main() {
	var joao, mes, carlos float64
	f.Scan(&carlos)
	joao = (carlos / 3)
	for joao < carlos {
		carlos *= 1.02
		joao *= 1.05
		mes++
	}
	f.Print(mes)
}
