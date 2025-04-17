package main

import "fmt"

func main() {
	var n, answ int
	fmt.Scan(&n)
	for i := 0; i < (n / 2); i++ {
		if i*(i+1)*(i+2) == n {
			answ = 1
			break
		} else {
			answ = 2
		}
	}
	if answ == 1 {
		fmt.Print("O número é triangular.")
	} else {
		fmt.Print("O número não é triangular.")
	}
}
