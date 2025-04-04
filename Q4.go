package main

import "fmt"

func main() {
	var salario, kw, custokw, costoconsumo, custodesconto float64
	fmt.Print("Digite o valor do salário: ")
	fmt.Scan(&salario)
	fmt.Print("Digite a quantidade de kW consumida: ")
	fmt.Scan(&kw)
	custokw = (0.7 / 100) * salario
	costoconsumo = kw * custokw
	custodesconto = costoconsumo * (90 / 100)
	fmt.Printf("O custo por kW é: R$ %.2f\nO custo total do consumo é: R$ %.2f\nO custo total do consumo com desconto é: R$ %.2f\n", custokw, costoconsumo, custodesconto)
}
