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