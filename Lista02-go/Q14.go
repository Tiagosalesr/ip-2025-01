package main

import (
	"fmt"
	"strings"
)

func main() {
	var precoBase, precoFinal float64
	var resposta string

	arCondicionado := 1750.0
	pinturaMetalica := 800.0
	vidroEletrico := 1200.0
	direcaoHidraulica := 2000.0

	fmt.Print("Digite o valor de fábrica do carro: R$ ")
	fmt.Scanln(&precoBase)

	precoFinal = precoBase

	fmt.Print("Deseja adicionar Ar Condicionado? (s/n): ")
	fmt.Scanln(&resposta)
	if strings.ToLower(resposta) == "s" {
		precoFinal += arCondicionado
	}

	fmt.Print("Deseja adicionar Pintura Metálica? (s/n): ")
	fmt.Scanln(&resposta)
	if strings.ToLower(resposta) == "s" {
		precoFinal += pinturaMetalica
	}

	fmt.Print("Deseja adicionar Vidro Elétrico? (s/n): ")
	fmt.Scanln(&resposta)
	if strings.ToLower(resposta) == "s" {
		precoFinal += vidroEletrico
	}

	fmt.Print("Deseja adicionar Direção Hidráulica? (s/n): ")
	fmt.Scanln(&resposta)
	if strings.ToLower(resposta) == "s" {
		precoFinal += direcaoHidraulica
	}

	fmt.Printf("\nO preço final do carro com os opcionais escolhidos é: R$ %.2f\n", precoFinal)
}
