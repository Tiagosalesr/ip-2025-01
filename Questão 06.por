programa {
  inclua biblioteca Matematica --> mat
  inteiro c
  real f, cel, n1
  funcao inicio() {
    escreva("Digite o número de conversões a serem feitas: ")
    leia(n1)
    real vetor[n1]
    para(c = 0; c < n1; c = c +1)
   {  escreva("Digite a temperatura em fahrenheit: ")
   leia(f) 
   vetor[c] = (5/9) * ( f - 32 )}
    para(c = 0; c < n1; c = c + 1) {
      escreva("Temperatura ", c + 1, ": ", mat.arredondar(vetor[c],2), "\n")
  }
  
  }
}
