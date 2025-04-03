package main
import "fmt"
func main() {
	var nota1, nota2, nota3 float64
	fmt.Scanln(&nota1)
	fmt.Scanln(&nota2)
	fmt.Scanln(&nota3)
	media := (nota1 + nota2 + nota3) / 3
	fmt.Printf("MEDIA = %.2f\n", media)
	if media >= 6 {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
	}
