// Package winui expõe caixas de texto e de mensagem nativas do Windows
// através do PowerShell, evitando repetir o mesmo script em cada aplicativo.
package winui

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// ErrCancelado indica que o usuário fechou a caixa de diálogo sem informar valor.
var ErrCancelado = errors.New("diálogo cancelado pelo usuário")

const carregaAssemblies = `
[void][System.Reflection.Assembly]::LoadWithPartialName('Microsoft.VisualBasic')
[void][System.Reflection.Assembly]::LoadWithPartialName('System.Windows.Forms')
`

// InputBox mostra uma caixa de texto e devolve o que foi digitado.
func InputBox(prompt, titulo string) (string, error) {
	script := fmt.Sprintf(
		"%s\nWrite-Output ([Microsoft.VisualBasic.Interaction]::InputBox(%s, %s))",
		carregaAssemblies, aspas(prompt), aspas(titulo),
	)
	saida, err := executar(script)
	if err != nil {
		return "", err
	}
	texto := strings.TrimSpace(saida)
	if texto == "" {
		return "", ErrCancelado
	}
	return texto, nil
}

// MessageBox mostra uma janela de mensagem simples.
func MessageBox(mensagem, titulo string) error {
	script := fmt.Sprintf(
		"%s\n[void][System.Windows.Forms.MessageBox]::Show(%s, %s)",
		carregaAssemblies, aspas(mensagem), aspas(titulo),
	)
	_, err := executar(script)
	return err
}

// PerguntarFloat repete a caixa de texto até receber um número válido.
func PerguntarFloat(prompt, titulo string) (float64, error) {
	for {
		texto, err := InputBox(prompt, titulo)
		if err != nil {
			return 0, err
		}
		valor, err := strconv.ParseFloat(strings.Replace(texto, ",", ".", 1), 64)
		if err == nil {
			return valor, nil
		}
		if err := MessageBox("Valor inválido, tente novamente.", titulo); err != nil {
			return 0, err
		}
	}
}

func executar(script string) (string, error) {
	cmd := exec.Command("powershell", "-NoProfile", "-Command", script)
	ocultarJanela(cmd)
	saida, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("powershell: %w", err)
	}
	return string(saida), nil
}

// aspas transforma o texto em um literal PowerShell entre apóstrofos.
func aspas(texto string) string {
	return "'" + strings.ReplaceAll(texto, "'", "''") + "'"
}
