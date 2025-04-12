package main

import "fmt"

func main() {
	var (
		id                              int
		n1, n2, n3, mediaex, mediaaprov float64
		conceito                        string
	)
	fmt.Print("Digite o ID do aluno, as três notas das avaliações e a média de exercícios: ")
	fmt.Scan(&id, &n1, &n2, &n3, &mediaex)
	mediaaprov = (n1 + (n2 * 2) + (n3 * 3) + mediaex) / 7
	if mediaaprov >= 9 && mediaaprov <= 10 {
		conceito = "A"
	} else if mediaaprov >= 7.5 && mediaaprov < 9 {
		conceito = "B"
	} else if mediaaprov >= 6 && mediaaprov < 7.5 {
		conceito = "C"
	} else if mediaaprov >= 4 && mediaaprov < 6 {
		conceito = "D"
	} else if mediaaprov < 4 {
		conceito = "E"
	}
	fmt.Printf("O aluno nmr. %.2d tirou %.2f, %.2f e %.2f respectivamente nas três avaliaçõese %.2f na média de seus exercícios. Conseguindo uma média final de %.2f e conceito %s.\n", id, n1, n2, n3, mediaex, mediaaprov, conceito)
	if conceito == "A" || conceito == "B" || conceito == "C" {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}
