programa {
  inclua biblioteca Matematica --> mat
  real n1, n2, p, g, a, c, k, r1
  funcao inicio() {
   
 escreva("Digite a quantidade de casos de teste = ")
  leia(n1)
 real vetor[n1]
para (k = 0; k<= n1-1; k++) { escreva("Número de pessoas que compraram ingresso = ")
  leia(n2)
  escreva("Porcentagem de pessoas que compraram ingresso na categoria Popular = ")
  leia(p)
  escreva("Porcentagem de pessoas que compraram ingresso na categoria Geral = ")
  leia(g)
  escreva("Porcentagem de pessoas que compraram ingresso na categoria Arquibancada = ")
  leia(a)
  escreva("Porcentagem de pessoas que compraram ingresso na categoria Cadeiras = ")
  leia(c)
  vetor[k] = (p / 100) * (n2 * 1) + (g / 100) * (n2 * 5) + (a / 100) * (n2 * 10) + (c / 100) * (n2 * 20)}
  para (k = 0; k<= n1-1; k++) {escreva("A RENDA DO JOGO ", k + 1, " E = ", mat.arredondar(vetor[k],2),"\n")
  } 
}
}