package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"meuapp/imc"
)

func main() {
	meuApp := app.New()
	minhaJanela := meuApp.NewWindow("Calculadora de IMC")

	const peso = 100.00
	const altura = 1.75

	titulo := widget.NewLabel("-- Exemplo de classificação do IMC --")
	infoPeso := widget.NewLabel("Peso: " + imc.Formatar(peso) + " kg")
	infoAltura := widget.NewLabel("Altura: 1.75 m")
	resultadoIMC := widget.NewLabel("IMC: --")
	resultadoClassif := widget.NewLabel("Classificação: --")

	botaoCalcular := widget.NewButton("Calcular IMC", func() {
		valor, err := imc.Calcular(peso, altura)
		if err != nil {
			resultadoIMC.SetText("IMC: --")
			resultadoClassif.SetText("Classificação: " + err.Error())
			return
		}
		resultadoIMC.SetText("IMC: " + imc.Formatar(valor))
		resultadoClassif.SetText("Classificação: " + imc.Classificar(valor))
	})

	conteudo := container.NewVBox(titulo, infoPeso, infoAltura, botaoCalcular, resultadoIMC, resultadoClassif)
	minhaJanela.SetContent(conteudo)
	minhaJanela.ShowAndRun()
}
