package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Script limpo que usa as caixas de texto nativas do Windows
	psScript := `
	$ErrorActionPreference = 'Stop'
	[void][System.Reflection.Assembly]::LoadWithPartialName('Microsoft.VisualBasic')
	[void][System.Reflection.Assembly]::LoadWithPartialName('System.Windows.Forms')
	$p = [Microsoft.VisualBasic.Interaction]::InputBox('Digite o seu Peso em kg (Ex: 80):', 'Peso')
	if ($p -eq '') { exit }
	$a = [Microsoft.VisualBasic.Interaction]::InputBox('Digite a sua Altura em metros (Ex: 1.75):', 'Altura')
	if ($a -eq '') { exit }
	$peso = 0.0
	$altura = 0.0
	if (-not [double]::TryParse($p, [ref]$peso) -or $peso -le 0) {
		[Console]::Error.WriteLine('Peso invalido: ' + $p)
		exit 1
	}
	if (-not [double]::TryParse($a, [ref]$altura) -or $altura -le 0) {
		[Console]::Error.WriteLine('Altura invalida: ' + $a)
		exit 1
	}
	$imc = $peso / ($altura * $altura)
	$imcF = '{0:N2}' -f $imc
	[void][System.Windows.Forms.MessageBox]::Show('Seu IMC é: ' + $imcF, 'Resultado')
	`
	if err := runIMC(psScript); err != nil {
		fmt.Fprintln(os.Stderr, "erro ao calcular o IMC:", err)
		os.Exit(1)
	}
}

func runIMC(psScript string) error {
	var stderr bytes.Buffer
	cmd := exec.Command("powershell", "-Command", psScript)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if msg := bytes.TrimSpace(stderr.Bytes()); len(msg) > 0 {
			return fmt.Errorf("powershell: %w: %s", err, msg)
		}
		return fmt.Errorf("powershell: %w", err)
	}
	return nil
}
