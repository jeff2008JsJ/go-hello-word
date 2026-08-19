package main

import "fmt"

// calcularIMC devolve o IMC para o peso em kg e a altura em metros.
// Retorna erro quando os valores não são fisicamente válidos.
func calcularIMC(peso, altura float64) (float64, error) {
	if peso <= 0 {
		return 0, fmt.Errorf("peso inválido: %v", peso)
	}
	if altura <= 0 {
		return 0, fmt.Errorf("altura inválida: %v", altura)
	}
	return peso / (altura * altura), nil
}

// formatarIMC formata o IMC com duas casas decimais.
func formatarIMC(imc float64) string {
	return fmt.Sprintf("%.2f", imc)
}

// scriptIMC monta o script PowerShell que pede peso e altura ao usuário
// e mostra o IMC calculado em caixas de diálogo nativas do Windows.
func scriptIMC() string {
	return `
	[void][System.Reflection.Assembly]::LoadWithPartialName('Microsoft.VisualBasic')
	[void][System.Reflection.Assembly]::LoadWithPartialName('System.Windows.Forms')
	$p = [Microsoft.VisualBasic.Interaction]::InputBox('Digite o seu Peso em kg (Ex: 80):', 'Peso')
	if ($p -eq '') { exit }
	$a = [Microsoft.VisualBasic.Interaction]::InputBox('Digite a sua Altura em metros (Ex: 1.75):', 'Altura')
	if ($a -eq '') { exit }
	$imc = [double]$p / ([double]$a * [double]$a)
	$imcF = '{0:N2}' -f $imc
	[System.Windows.Forms.MessageBox]::Show('Seu IMC é: ' + $imcF, 'Resultado')
	`
}
