programa {
   inclua biblioteca Matematica --> mat
  funcao inicio() {
    real a, b, c, x
    escreva("Digite a primeira nota = ")
    leia(a)
    escreva("Digite a segunda nota = ")
    leia(b)
    escreva("Digite a terceira nota = ")
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
