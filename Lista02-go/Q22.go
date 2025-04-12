package main

import "fmt"

func main() {
	var (
		matricula                                                          int
		qntdhoraex, salarioBruto, descinss, descimpderenda, salarioliquido float64
	)
	fmt.Println("Digite o número de matrícula e a quantidade de horas extra trabalhadas:")
	fmt.Scanln(&matricula, &qntdhoraex)
	salarioBruto = 788 + (qntdhoraex * 10)
	if salarioBruto > 1500 {
		descinss = 0.12 * salarioBruto
	} else {
		descinss = 0
	}
	if salarioBruto > 2000 {
		descimpderenda = 0.2 * salarioBruto
	} else {
		descimpderenda = 0
	}
	salarioliquido = salarioBruto - (descinss + descimpderenda)
	fmt.Println("O número da matrícula é:", matricula)
	fmt.Printf("O salário líquido é de: R$ %.2f", salarioliquido)
}
