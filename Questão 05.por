programa {
  inclua biblioteca Matematica --> mat
  real conta, consumo, valor
  caracter tipo
  funcao inicio() {
    escreva("Número da Conta: ")
    leia(conta)
    escreva("Digite o consumo em metros cúbicos: ")
    leia(consumo)
    escreva("Digite o tipo da conta: ")
    leia(tipo)
     se (tipo == "R" ou tipo == "r")
    {
      valor = (0.05 * consumo) + 5
    } senao se (tipo == "C" ou tipo == "c" e consumo <= 80) 
    { 
      valor = 500
    } senao se( tipo == "C" ou tipo == "c" e consumo > 80)
    {
      valor = 500 + ((consumo - 80) * 0.25) 
    } senao se ( tipo == "I" ou tipo == "i" e consumo <= 100 ) 
    {
      valor = 800
    } senao se (tipo == "I" ou tipo == "i" e consumo > 100)
    {
      valor = 800 + ((consumo - 100) * 0.04)
    }
    escreva("CONTA = ",mat.arredondar(conta,2), "\n")
    escreva("VALOR DA CONTA = ",mat.arredondar(valor,2),"\n")
  }
}
