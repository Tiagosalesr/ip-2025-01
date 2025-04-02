package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3 int

	fmt.Scanln(&n1, &n2, &n3)

   if (n1 <= 9 && n1 >= 0 && n2 <= 9 && n2 >= 0 && n3 <= 9 && n3 >= 0 ) {
	numeroFinal := (n1 * 100) + (n2 * 10) 
	fmt.Println(numeroFinal,n3,", ")
	n4 := numeroFinal + n3
	fmt.Println(n4 * n4)
   } else {
   fmt.Println("DIGITO INVALIDO")
   }


}
