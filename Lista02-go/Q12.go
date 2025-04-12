package main

import "fmt"

func main() {
	var (
		idade int
		class string
	)
	for true {
		fmt.Print("Digite a idade: ")
		fmt.Scan(&idade)
		if idade < 0 {
			fmt.Println("Idade inválida")
		} else {
			break
		}
	}
	if idade <= 2 {
		class = ("Recém-nascido")
	} else if idade > 2 && idade <= 11 {
		class = ("Criança")
	} else if idade > 11 && idade <= 19 {
		class = ("Adolescente")
	} else if idade > 19 && idade <= 55 {
		class = ("Adulto")
	} else {
		class = ("Idoso")
	}
	fmt.Print(class)
}
