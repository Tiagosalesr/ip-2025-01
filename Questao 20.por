programa {
  funcao inicio() {
    real h, min, seg
    leia(h)
    leia(min)
    leia(seg)
    h = h * 3600
    min = min * 60
    escreva("O TEMPO EM SEGUNDOS E = ", min + h + seg)
  }
}
