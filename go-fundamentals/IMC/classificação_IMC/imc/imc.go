// Package imc calcula e classifica o Índice de Massa Corporal.
package imc

import "fmt"

// Calcular devolve o IMC para o peso em kg e a altura em metros.
// Retorna erro quando os valores não são fisicamente válidos.
func Calcular(peso, altura float64) (float64, error) {
	if peso <= 0 {
		return 0, fmt.Errorf("peso inválido: %v", peso)
	}
	if altura <= 0 {
		return 0, fmt.Errorf("altura inválida: %v", altura)
	}
	return peso / (altura * altura), nil
}

// Classificar devolve a classificação do IMC segundo as faixas da OMS.
func Classificar(imc float64) string {
	switch {
	case imc < 18.5:
		return "Abaixo do peso"
	case imc < 25:
		return "Peso normal"
	case imc < 30:
		return "Sobrepeso"
	case imc < 35:
		return "Obesidade grau I"
	case imc < 40:
		return "Obesidade grau II"
	default:
		return "Obesidade grau III"
	}
}

// Formatar formata o IMC com duas casas decimais.
func Formatar(imc float64) string {
	return fmt.Sprintf("%.2f", imc)
}
