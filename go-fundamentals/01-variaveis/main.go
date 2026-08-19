package main

import (
	"errors"
	"log"
	"os"

	"github.com/jeff2008JsJ/go-hello-word/pkg/imc"
	"github.com/jeff2008JsJ/go-hello-word/pkg/winui"
)

func main() {
	peso, err := winui.PerguntarFloat("Digite o seu Peso em kg (Ex: 80):", "Peso")
	if err != nil {
		encerrar(err)
	}
	altura, err := winui.PerguntarFloat("Digite a sua Altura em metros (Ex: 1.75):", "Altura")
	if err != nil {
		encerrar(err)
	}

	resumo, err := imc.Resumo(peso, altura)
	if err != nil {
		if err := winui.MessageBox(err.Error(), "Resultado"); err != nil {
			log.Fatal(err)
		}
		return
	}

	if err := winui.MessageBox(resumo, "Resultado"); err != nil {
		log.Fatal(err)
	}
}

func encerrar(err error) {
	if errors.Is(err, winui.ErrCancelado) {
		os.Exit(0)
	}
	log.Fatal(err)
}
