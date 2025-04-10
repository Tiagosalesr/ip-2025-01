package main

import "fmt"

func main(){
	var valor, codigo, valorcdesconto float64
	for true {

	fmt.Scan(&valor,&codigo)
	if codigo < 1 || codigo > 4 {
		fmt.Print("Código inválido")
	}else {
		break
	}
}
if codigo == 1 {
	valorcdesconto = valor * 0.9
	fmt.Printf("O valor a se pagar é : R$ %.2f \n",valorcdesconto)
}else if codigo == 2{
	valorcdesconto = valor * 0.95
	fmt.Printf("O valor a se pagar é : R$ %.2f \n",valorcdesconto)
}else if codigo == 3{
	valorcdesconto = valor 
	fmt.Printf("O valor a se pagar é R$ %.2f em duas vezes sem juros.\n",valorcdesconto)
}else {
	valorcdesconto = valor * 1.1
	fmt.Print("O valor a se pagar é R$ %.2f em 3 vezem com 10'%' de juros. \n",valorcdesconto)
}

}
