programa {
  inteiro x, y, i
  funcao inicio() {
    escreva("Digite o primeiro número: ")
    leia(x)
    escreva("Digite o segundo número: ")
    leia(y)
    se( x%2 == 0) {
      para(i=0; i < y; i = i + 1 ) {escreva(x + (2 * i), " ")}
      } senao {escreva("O PRIMEIRO NUMERO NAO E PAR")}
  }
}
