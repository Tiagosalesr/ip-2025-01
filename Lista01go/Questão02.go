package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	valoresIngressos := map[string]float64{
		"Popular":     1.00,
		"Geral":       5.00,
		"Arquibancada": 10.00,
		"Cadeiras":    20.00,
	}


	for i := 1; i <= t; i++ {
		
		var totalPessoas, percPopular, percGeral, percArquibancada, percCadeiras int
		fmt.Scan(&totalPessoas, &percPopular, &percGeral, &percArquibancada, &percCadeiras)

		
		numPopular := float64(totalPessoas * percPopular) / 100
		numGeral := float64(totalPessoas * percGeral) / 100
		numArquibancada := float64(totalPessoas * percArquibancada) / 100
		numCadeiras := float64(totalPessoas * percCadeiras) / 100

		
		renda := (numPopular * valoresIngressos["Popular"]) +
			(numGeral * valoresIngressos["Geral"]) +
			(numArquibancada * valoresIngressos["Arquibancada"]) +
			(numCadeiras * valoresIngressos["Cadeiras"])

		
		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, renda)
	}
}
