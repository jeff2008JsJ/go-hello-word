// Package imc concentra o cálculo, a formatação e a classificação do
// Índice de Massa Corporal usados pelos aplicativos deste repositório.
package imc

import (
	"errors"
	"fmt"
)

// ErrAlturaInvalida é retornado quando a altura informada não é positiva.
var ErrAlturaInvalida = errors.New("altura deve ser maior que zero")

// ErrPesoInvalido é retornado quando o peso informado não é positivo.
var ErrPesoInvalido = errors.New("peso deve ser maior que zero")

// Calcular retorna o IMC para um peso em quilos e uma altura em metros.
func Calcular(pesoKg, alturaM float64) (float64, error) {
	if pesoKg <= 0 {
		return 0, ErrPesoInvalido
	}
	if alturaM <= 0 {
		return 0, ErrAlturaInvalida
	}
	return pesoKg / (alturaM * alturaM), nil
}

// Classificar retorna a faixa da OMS correspondente ao IMC.
func Classificar(valor float64) string {
	switch {
	case valor < 18.5:
		return "Abaixo do peso"
	case valor < 25:
		return "Peso normal"
	case valor < 30:
		return "Sobrepeso"
	case valor < 35:
		return "Obesidade grau I"
	case valor < 40:
		return "Obesidade grau II"
	default:
		return "Obesidade grau III"
	}
}

// Formatar devolve o IMC com duas casas decimais.
func Formatar(valor float64) string {
	return fmt.Sprintf("%.2f", valor)
}

// FormatarPeso devolve o peso com duas casas decimais e unidade.
func FormatarPeso(pesoKg float64) string {
	return fmt.Sprintf("%.2f kg", pesoKg)
}

// FormatarAltura devolve a altura com duas casas decimais e unidade.
func FormatarAltura(alturaM float64) string {
	return fmt.Sprintf("%.2f m", alturaM)
}

// Resumo descreve o IMC e a sua classificação em uma única linha.
func Resumo(pesoKg, alturaM float64) (string, error) {
	valor, err := Calcular(pesoKg, alturaM)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("IMC: %s (%s)", Formatar(valor), Classificar(valor)), nil
}
