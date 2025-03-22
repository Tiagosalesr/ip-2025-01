programa {
  real n1,n2,n3,k,l,n4
  funcao inicio() {
    escreva("Digite o primeiro número = ")
    leia(n1)
    escreva("Digite o segundo número = ")
    leia(n2)
    escreva("Digite o terceiro número = ")
    leia(n3)
    se (n1 >= 10 ou n2 >= 10 ou n3 >=10) escreva("DIGITO INVALIDO")
    senao
    k = n1*100
    l = n2*10
    escreva(k + l + n3)
    n4 = k + l + n3
    escreva(", ", n4*n4)
  }
}
