package main

import (
	"fmt"
)

func main() {
	soma := 0
	for i := 1; i <= 20; i++ {
		fmt.Print(i, " ")
		soma += i
	}
	fmt.Print("\n", soma)
}
