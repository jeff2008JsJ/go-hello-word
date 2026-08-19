package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/jeff2008JsJ/go-hello-word/pkg/imc"
)

func main() {
	const (
		peso   = 100.00
		altura = 1.75
	)

	meuApp := app.New()
	minhaJanela := meuApp.NewWindow("Calculadora de IMC")

	titulo := widget.NewLabel("-- Exemplo de classificação do IMC --")
	infoPeso := widget.NewLabel("Peso: " + imc.FormatarPeso(peso))
	infoAltura := widget.NewLabel("Altura: " + imc.FormatarAltura(altura))
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
