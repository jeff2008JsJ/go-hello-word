package main

import (
"fmt"
"fyne.io/fyne/v2/app"
"fyne.io/fyne/v2/container"
"fyne.io/fyne/v2/widget"
)

func main() {
meuApp := app.New()
minhaJanela := meuApp.NewWindow("Calculadora de IMC")

titulo := widget.NewLabel("-- Exemplo de classificação do IMC --")
infoPeso := widget.NewLabel("Peso: 100.00 kg")
infoAltura := widget.NewLabel("Altura: 1.75 m")
resultadoIMC := widget.NewLabel("IMC: --")
resultadoClassif := widget.NewLabel("Classificação: --")

botaoCalcular := widget.NewButton("Calcular IMC", func() {
peso := 100.00
altura := 1.75
imc := peso / (altura * altura)
resultadoIMC.SetText(fmt.Sprintf("IMC: %.2f", imc))
resultadoClassif.SetText("Classificação: Obesidade grau I")
})

conteudo := container.NewVBox(titulo, infoPeso, infoAltura, botaoCalcular, resultadoIMC, resultadoClassif)
minhaJanela.SetContent(conteudo)
minhaJanela.ShowAndRun()
}
