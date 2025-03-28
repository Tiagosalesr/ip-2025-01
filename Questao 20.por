programa {
  funcao inicio() {
    real h, min, seg
    escreva("Digite a hora: ")
    leia(h)
    escreva("Digite os minutos: ")
    leia(min)
    escreva("Digite os segundos: ")
    leia(seg)
    h = h * 3600
    min = min * 60
    escreva("O TEMPO EM SEGUNDOS E = ", min + h + seg)
  }
}
