programa {
  funcao inicio() {
    inteiro i, a1, razao, numeroElementos, soma
    soma = 0
    escreva("Digite o primeiro termo da P.A.: ")
    leia(a1)
    escreva("Digite a razão da P.A.: ")
    leia(razao)
    escreva("Digite o número de elementos da P.A.: ")
    leia(numeroElementos)  
   real vetor[numeroElementos]
    para (i=0; i < numeroElementos; i++) {vetor[i] = (a1 + (razao * i))}
  
  para (i=0; i < numeroElementos; i++) {soma = soma + vetor[i]}
  escreva("Soma dos itens da P.A. = ", soma)
}
}

