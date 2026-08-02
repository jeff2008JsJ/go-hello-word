package main

import (
	"os/exec"
	"syscall"
)

func main() {
	// Script limpo que usa as caixas de texto nativas do Windows
	psScript := `
	[void][System.Reflection.Assembly]::LoadWithPartialName('Microsoft.VisualBasic')
	$p = [Microsoft.VisualBasic.Interaction]::InputBox('Digite o seu Peso em kg (Ex: 80):', 'Peso')
	if ($p -eq '') { exit }
	$a = [Microsoft.VisualBasic.Interaction]::InputBox('Digite a sua Altura em metros (Ex: 1.75):', 'Altura')
	if ($a -eq '') { exit }
	$imc = [double]$p / ([double]$a * [double]$a)
	$imcF = '{0:N2}' -f $imc
	[System.Windows.Forms.MessageBox]::Show('Seu IMC é: ' + $imcF, 'Resultado')
	`
	cmd := exec.Command("powershell", "-Command", psScript)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()
}
