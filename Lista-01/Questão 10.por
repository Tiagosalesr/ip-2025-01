programa {
  inclua biblioteca Matematica --> mat
  real a,b,c,d,det
  funcao inicio() {
    escreva("Defina a: ")
    leia(a)
    escreva("Defina b: ")
    leia(b)
    escreva("Defina c: ")
    leia(c)
    escreva("Defina d: ")
    leia(d)
    det = (a * d) - (b * c)
    escreva("O VALOR DO DETERMINANTE É = ",mat.arredondar(det,2),"\n")
  }
}
