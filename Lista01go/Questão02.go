package main

import (
	"fmt"
)

func main() {
	var (
		jogos int
		pubtot, popular, geral, arquibancada, cadeiras float32
	)
	receita := [4]float32{}
	receitatot := [3]float32{}

	fmt.Print("São quantos jogos? ")
	fmt.Scanln(&jogos)


	for i:= 0; i<= jogos - 1; i++ {

	fmt.Print("Digite: O número de pessoas que compraram ingresso para o jogo N.", i+1,
		  ", a percentagem de pessoas que compraram ingresso na categoria Popular do jogo N.", i+1,
		  ", a percentagem de pessoas que compraram ingresso na categoria Geral do jogo N.", i+1,
		  ", a percentagem de pessoas que compraram ingresso na categoria Arquibancada do jogo N.", i+1 , 
		  " e ", "a percentagem de pessoas que compraram ingresso na categoria Cadeiras do jogo N.", i+1,": " )
	fmt.Scanln(&pubtot,&popular,&geral,&arquibancada,&cadeiras)

	receita[0] = pubtot * (popular/100) * 1
       	receita[1] = pubtot * (geral/100) * 5
       	receita[2] = pubtot * (arquibancada/100) * 10
       	receita[3] = pubtot * (cadeiras/100) * 20

		receitatot[i] = receita[0] + receita[1] + receita[2] + receita[3]
}
for i:= 0; i < jogos; i++ {
	fmt.Printf("A RENDA DO JOGO N. %d É DE: R$ %.2f\n", i+1, receitatot[i])
}
}
