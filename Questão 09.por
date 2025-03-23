programa {
  inclua biblioteca Matematica --> mat
  real delta, a, b, c
  funcao inicio() {
    leia(a)
    leia(b)
    leia(c)
    delta = (b * b) - (4 * a * c)
    escreva("O VALOR DE DELTA E = ",mat.arredondar(delta,2),"\n")
  }
}
