package main

import (
	"math"
	"os/exec"
	"strings"
	"testing"
)

func TestCalcularIMC(t *testing.T) {
	casos := []struct {
		nome   string
		peso   float64
		altura float64
		quer   float64
	}{
		{"peso e altura típicos", 80, 1.80, 24.691358},
		{"altura de 1 metro", 70, 1, 70},
		{"obesidade grau I", 100, 1.75, 32.653061},
	}

	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			got, err := calcularIMC(c.peso, c.altura)
			if err != nil {
				t.Fatalf("erro inesperado: %v", err)
			}
			if math.Abs(got-c.quer) > 1e-6 {
				t.Errorf("calcularIMC(%v, %v) = %v, quer %v", c.peso, c.altura, got, c.quer)
			}
		})
	}
}

func TestCalcularIMCInvalido(t *testing.T) {
	casos := []struct {
		nome   string
		peso   float64
		altura float64
	}{
		{"peso zero", 0, 1.75},
		{"peso negativo", -5, 1.75},
		{"altura zero", 80, 0},
		{"altura negativa", 80, -1.75},
	}

	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			imc, err := calcularIMC(c.peso, c.altura)
			if err == nil {
				t.Fatalf("calcularIMC(%v, %v) = %v, queria erro", c.peso, c.altura, imc)
			}
			if imc != 0 {
				t.Errorf("IMC = %v em caso de erro, quer 0", imc)
			}
		})
	}
}

func TestFormatarIMC(t *testing.T) {
	casos := []struct {
		imc  float64
		quer string
	}{
		{24.691358, "24.69"},
		{32.6530612, "32.65"},
		{18.5, "18.50"},
	}

	for _, c := range casos {
		if got := formatarIMC(c.imc); got != c.quer {
			t.Errorf("formatarIMC(%v) = %q, quer %q", c.imc, got, c.quer)
		}
	}
}

func TestScriptIMC(t *testing.T) {
	script := scriptIMC()

	for _, esperado := range []string{
		"Microsoft.VisualBasic",
		"System.Windows.Forms",
		"Peso",
		"Altura",
		"MessageBox",
	} {
		if !strings.Contains(script, esperado) {
			t.Errorf("script não contém %q", esperado)
		}
	}
}

func TestEsconderJanelaNaoFalha(t *testing.T) {
	cmd := exec.Command("powershell", "-Command", scriptIMC())
	esconderJanela(cmd)
	if len(cmd.Args) != 3 {
		t.Errorf("cmd.Args = %v, quer 3 argumentos", cmd.Args)
	}
}
