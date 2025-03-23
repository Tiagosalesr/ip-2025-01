programa {
  inclua biblioteca Matematica --> mat
  inteiro c
  real f, cel, n1
  funcao inicio() {
    escreva("Digite o número de conversões a serem feitas: ")
    leia(n1)
    faca {
    c=0 
    c = c + 1
    leia(f)
    cel = (5/9) * ( f - 32 )
    escreva(f, " FAHRENHEIT EQUIVALE A ", mat.arredondar(cel,2), " CELSIUS.","\n")}
    enquanto (c < n1)
  }
}
