package main

import "fmt"

func main() {
	var (
		dia, mes, ano int
		month         string
	)
	for true {
		fmt.Print("Digite o dia, mês e ano separados por um espaço: ")
		fmt.Scan(&dia, &mes, &ano)
		if dia < 1 || dia >= 32 || mes < 1 || mes > 12 {
			fmt.Println("Dia ou mês Inválido")
		} else {
			break
		}
	}
	if mes == 1 {
		month = ("Janeiro")
	} else if mes == 2 {
		month = ("Fevereiro")
	} else if mes == 3 {
		month = ("Março")
	} else if mes == 4 {
		month = ("Abril")
	} else if mes == 5 {
		month = ("Maio")
	} else if mes == 6 {
		month = ("Junho")
	} else if mes == 7 {
		month = ("Julho")
	} else if mes == 8 {
		month = ("Agosto")
	} else if mes == 9 {
		month = ("Setembro")
	} else if mes == 10 {
		month = ("Outubro")
	} else if mes == 11 {
		month = ("Novembro")
	} else {
		month = ("Dezembro")
	}
	fmt.Printf("%d de %s de %d", dia, month, ano)
}
