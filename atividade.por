programa {
   inclua biblioteca Matematica --> mat
  funcao inicio() {
    real a, b, c, x, y

    escreva("Digite o primeiro número = ")
    leia(a)

    escreva("Digite o segundo número = ")
    leia(b)

    escreva("Digite o terceiro número = ")
    leia(c)

    x = (a + b + c)/3 

    limpa()

    escreva("Média = ",mat.arredondar(x,2), "\n") 


    se(x >= 6) 
    {
      escreva("Aprovado\n")
    }

    senao{
      escreva("Reprovado\n")
    }

  }
}