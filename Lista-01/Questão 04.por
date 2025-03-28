programa {
  inclua biblioteca Matematica --> mat 
   real salario, kw, ckw, cc, ccd
  funcao inicio() {
   escreva("Valor do salário: ")
   leia(salario)
   escreva("Quantidade de kW: ")
   leia(kw)
   ckw = (0.7/100) * salario
   cc = kw * ckw
   ccd = (90/100) * cc
    escreva("Custo por kW: R$ ", mat.arredondar(ckw,2), "\n", 
    "Custo do consumo: R$ ", mat.arredondar(cc,2), "\n", 
    "Custo com desconto: R$ ", mat.arredondar(ccd,2)
    )
  }
}
