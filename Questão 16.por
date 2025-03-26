programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real salario, reajuste
    escreva("Escreva o salário atual: ")
    leia(salario)
    se( salario <= 300 ) { reajuste = (15/10) * salario} 
    senao {reajuste = (13/10) * salario }
   escreva("SALARIO COM REAJUSTE = ", mat.arredondar(reajuste,2))
  }
}
