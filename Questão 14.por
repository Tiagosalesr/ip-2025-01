programa {
  inclua biblioteca Matematica --> mat
  real a, h, v, ab, raiz
  funcao inicio() {
    escreva("DIGITE A ALTURA : ")
    leia(h)
    escreva("DIGITE O COMPRIMENTO DA ARESTA : ")
    leia(a)
    raiz = mat.raiz(3,2.0)
    ab = (3 * a * a * raiz ) / (2)
    v = (ab / 3) * (h)
    escreva("O VOLUME DA PIRAMIDE E = ",mat.arredondar(v,2), " METROS CUBICOS.","\n")
  }
}
