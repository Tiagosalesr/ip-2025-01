programa {
  inclua biblioteca Matematica --> mat
  inteiro horas
  real valor, n1, c
  funcao inicio() {
    c = 0
    escreva("Digite o número de horas: " )
    leia(horas)
    se(horas%3==0) 
    {valor = (horas / 3) * 10 }
    senao { faca { 
    horas = horas - 1 
    c = c + 1} 
    enquanto(horas%3!=0)
     valor = (horas/3) * (10) + (c * 5)
    }
  escreva("O VALOR A PAGAR E = ",mat.arredondar(valor,2),"\n")
  }
}
