programa {
  inclua biblioteca Matematica --> mat
  real nota
  cadeia conceito
  funcao inicio() {
    escreva("Digite a nota: ")
    leia(nota)
    se(nota <= 10 e  nota >=9 ) {conceito = "A"} 
    senao se(nota < 9 e nota >= 7.5) {conceito = "B"}
    senao se(nota < 7.5  e nota >= 6) {conceito = "C"}
    senao se(nota < 6  e nota >= 0) {conceito = "D"}
    escreva("NOTA = ", mat.arredondar(nota,2), " CONCEITO = ",conceito,"\n")
  }
}
