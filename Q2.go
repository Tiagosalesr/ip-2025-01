package main

import "fmt"

func main (){
	
	array10 :=[]int{}
	array5 := []int{}
	vetorR1 := []int{}
	vetorR2 := []int{}
	
	montagemArray(&array10,10)
	montagemArray(&array5,5)

	soma:=0

	for _, v := range array5{
		
		soma +=v
	}
		var n int
	for _, v2 := range array10{
		if v2%2==0 {
			n = soma + v2
			vetorR1 = append(vetorR1, n)
		} else if v2%2!=0 {
			n = soma + v2
			vetorR2 = append(vetorR2, n)
		}
	}
	fmt.Printf("O vetor resultante 01 é: %d\n", vetorR1)
	fmt.Printf("O vetor resultante 02 é: %d\n", vetorR2)
}

func montagemArray(list *[]int, indice int){
	fmt.Printf("vetor com índice %d, digite os valores: \n", indice)
	for i:= 1; i <= indice; i++ {
		var num int
		fmt.Scan(&num)
		*list = append(*list,num)
	}
}