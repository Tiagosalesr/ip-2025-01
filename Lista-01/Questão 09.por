programa {
  inclua biblioteca Matematica --> mat
  real delta, a, b, c
  funcao inicio() {
    escreva("Digite a: ")
    leia(a)
    escreva("Digite b: ")
    leia(b)
    escreva("Digite c: ")
    leia(c)
    delta = (b * b) - (4 * a * c)
    escreva("O VALOR DE DELTA E = ",mat.arredondar(delta,2),"\n")
  }
}
