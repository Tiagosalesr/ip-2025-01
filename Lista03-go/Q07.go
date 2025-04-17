package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, soma, qntd, media, somapar, mediapar, qntdpar, qntdimpar, percimpar float64
	var list []float64
	list = append(list, 30000)
	for {
		fmt.Print("Digite um número: ")
		fmt.Scan(&n)
		if n < 30000 || n > 30000 {
			list = append(list, n)
			soma += n + 30000
			qntd += 1
		} else {
			break
		}
	}
	qntd += 1
	media = soma / qntd

	sort.Float64s(list)

	l := len(list)

	min, max := list[0], list[l-1]

	for _, num := range list {

		if int(num)%2 == 0 {
			somapar += num
			qntdpar += 1
		} else {
			qntdimpar += 1
		}
	}
	mediapar = somapar / qntdpar
	percimpar = (qntdimpar / float64(l)) * 100
	fmt.Printf("A soma dos número digitados é = %.2f\nA quantidade de números digitados é = %d\nA média dos números é = %.2f\nO maior termo é %.2f e o menor %.2f\nA média dos números pares é = %.2f\nA porcentagem de números ímpares é de %.1f/100", soma, l, media, max, min, mediapar, percimpar)
}
