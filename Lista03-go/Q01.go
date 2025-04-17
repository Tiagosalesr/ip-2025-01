package main

import "fmt"

func main() {
	var (
		base, answ float64
		exp        int
	)
	for {
		fmt.Print("Escreva a base e o expoente da potência: ")
		fmt.Scan(&base, &exp)
		if exp < 0 {
			fmt.Println("Expoente negativo.")
		} else {
			break
		}
	}
	if exp == 0 {
		answ = 1
	} else {
		answ = 1
		for i := 0; i < exp; i++ {
			answ *= base
		}
	}
	fmt.Print(answ)
}
