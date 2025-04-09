package main

import(
	"fmt"
	"math"
)

func main(){
   var (
	forma int
    altura, raio, volumeforma, areaforma float64
)
	for true {
		fmt.Println("Digite o número corresponde à sua figura(1-cone,2-cilindro,3-esfera) e a altura e o raio dela:")
		fmt.Scan(&forma,&altura,&raio)

		
		if forma > 3 || forma < 1 {
			fmt.Println("Dígito inválido") 
			
		}else {
			break
		}
	} 
   if forma == 1 {
	volumeforma = (math.Pi * math.Pow(raio,2) * altura) / 3
	areaforma = math.Pi * raio * math.Sqrt(math.Pow(raio,2)+math.Pow(altura,2))
   }else if forma == 2 {
	volumeforma = (math.Pi * math.Pow(raio,2) * altura)
	areaforma = 2 * math.Pi * raio * altura
   }else {
	volumeforma = (4/3) * math.Pi * math.Pow(raio,3)
	areaforma = 4 * math.Pi * math.Pow(raio,2)
   }
   fmt.Printf("O volume da figura é: %.f \n A área da figura é: %.f \n", volumeforma, areaforma)
}
