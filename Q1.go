package main

import "fmt"

func main (){
	var (
		num int
		array = []int{}
	)
	for i:=1;i<=10;i++ {
		fmt.Printf("Digite o %dº termo: ",i)
		fmt.Scan(&num)
		array = append(array,num)
	} 
	encontrou := false
	for i, v := range array {
		if v > 50 {
			fmt.Printf("O número %d da posião %dº é maior que cinquenta.\n",v,i)
			encontrou = true
		}
	}
	if !encontrou {
		fmt.Println("Não há números maiores que cinquenta.")
	}
}