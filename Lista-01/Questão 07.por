programa {
  inclua biblioteca Matematica --> mat
  real cel, fahr, pol, mm
  funcao inicio() {
    escreva("Digite a temperatura em fahrenheit: ")
    leia(fahr)
    cel = (5/9) * (fahr - 32)
    escreva("Digite o volume de chuva em polegadas: ")
    leia(pol)
    mm = pol*25.4
    escreva("O VALOR EM CELSIUS = ",mat.arredondar(cel,2), "\n", "A QUANTIDADE DE CHUVA É = ",mat.arredondar(mm,2), "\n")
  }
}
