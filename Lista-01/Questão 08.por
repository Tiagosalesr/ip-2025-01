programa {
  inclua biblioteca Matematica -->mat
  real at, al, ac, pi, r, a, valor
  funcao inicio() {
    escreva("Digite o raio: ")
    leia(r)
    escreva("Digite a altura: ")
    leia(a)
    pi = 3.14159
    al = (a * 2 * pi * r)
    ac = (pi * r * r)
    at =  (2 * ac) + al
    valor = at * 100
    escreva("O VALOR DO CUSTO E = ", "R$ ",mat.arredondar(valor,2))
  }
}
