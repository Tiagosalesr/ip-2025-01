package main

import "fmt"

func main() {
	var regiao, viagem, valor int
	fmt.Println("Digite a região de destino e se a viagem é somente de ida ou de ida e volta (separados por um espaço):")
	fmt.Scan(&regiao, &viagem)
	if regiao == 1 && viagem == 1 {
		valor = 500
	} else if regiao == 1 && viagem == 2 {
		valor = 900
	} else if regiao == 2 && viagem == 1 {
		valor = 350
	} else if regiao == 2 && viagem == 2 {
		valor = 650
	} else if regiao == 3 && viagem == 1 {
		valor = 350
	} else if regiao == 3 && viagem == 2 {
		valor = 600
	} else if regiao == 4 && viagem == 1 {
		valor = 300
	} else if regiao == 4 && viagem == 2 {
		valor = 550
	} else {
		valor = 0
	}
	if valor > 0 {
		fmt.Println("O valor da passagem é:", valor)
	} else {
		fmt.Println("Região ou tipo de viagem inválido.")
	}
}
