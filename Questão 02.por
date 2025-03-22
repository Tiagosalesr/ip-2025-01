programa {
  inclua biblioteca Matematica --> mat
  real n1, n2, p, g, a, c, k, r1
  funcao inicio() {
   k = 0
  escreva("Digite a quantidade de casos de teste = ")
  leia(n1)
  faca
  { escreva("Número de pessoas que compraram ingresso = ")
  leia(n2)
  escreva("Porcentagem de pessoas que compraram ingresso na categoria Popular = ")
  leia(p)
  escreva("Porcentagem de pessoas que compraram ingresso na categoria Geral = ")
  leia(g)
  escreva("Porcentagem de pessoas que compraram ingresso na categoria Arquibancada = ")
  leia(a)
  escreva("Porcentagem de pessoas que compraram ingresso na categoria Cadeiras = ")
  leia(c)
  k=k+1
  r1 = (p / 100) * (n2 * 1) + (g / 100) * (n2 * 5) + (a / 100) * (n2 * 10) + (c / 100) * (n2 * 20)
 escreva("A RENDA DO JOGO ", k, " E = ", mat.arredondar(r1,2),"\n")
  } enquanto (k <= n1-1)
}
}