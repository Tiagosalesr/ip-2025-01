package main

import "fmt"

func main() {
	var (
		idade, altura, peso, somaaltura, porc40menos         float64
		nmr, fiftymore, qntdaltura, qntd40menos, totalpessoa int
	)
	for {
		fmt.Print("Digite a idade o peso e a altura da pessoa e se deseja continuar inserindo dados(digite '1') ou não (digite um número diferente de '1'): ")
		fmt.Scan(&idade, &peso, &altura, &nmr)
		if nmr == 1 {
			if idade > 50 {
				fiftymore++
			}
			if idade > 10 && idade < 20 {
				somaaltura += altura
				qntdaltura++
			}
			if peso < 40 {
				qntd40menos++
			}
			totalpessoa++
		} else {
			break
		}
	}
	mediaaltura := somaaltura / float64(qntdaltura)
	porc40menos = (float64(qntd40menos / totalpessoa)) * 100
	fmt.Printf("Há %d com mais de 50 anos, a média de altura de pessoas entre 10 e 20 anos é %f e %.1f/100 das pessoas pesam menos de 40 Kg.", fiftymore, mediaaltura, porc40menos)
}
