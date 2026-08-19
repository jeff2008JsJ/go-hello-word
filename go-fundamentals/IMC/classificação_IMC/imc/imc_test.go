package imc

import (
	"math"
	"testing"
)

func TestCalcular(t *testing.T) {
	casos := []struct {
		nome   string
		peso   float64
		altura float64
		quer   float64
	}{
		{"exemplo do app", 100, 1.75, 32.6530612},
		{"peso normal", 70, 1.80, 21.6049382},
		{"altura de 1 metro", 50, 1, 50},
	}

	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			got, err := Calcular(c.peso, c.altura)
			if err != nil {
				t.Fatalf("erro inesperado: %v", err)
			}
			if math.Abs(got-c.quer) > 1e-6 {
				t.Errorf("Calcular(%v, %v) = %v, quer %v", c.peso, c.altura, got, c.quer)
			}
		})
	}
}

func TestCalcularInvalido(t *testing.T) {
	casos := []struct {
		nome   string
		peso   float64
		altura float64
	}{
		{"peso zero", 0, 1.75},
		{"peso negativo", -80, 1.75},
		{"altura zero", 80, 0},
		{"altura negativa", 80, -1.75},
	}

	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			got, err := Calcular(c.peso, c.altura)
			if err == nil {
				t.Fatalf("Calcular(%v, %v) = %v, queria erro", c.peso, c.altura, got)
			}
			if got != 0 {
				t.Errorf("IMC = %v em caso de erro, quer 0", got)
			}
		})
	}
}

func TestClassificar(t *testing.T) {
	casos := []struct {
		imc  float64
		quer string
	}{
		{16, "Abaixo do peso"},
		{18.49, "Abaixo do peso"},
		{18.5, "Peso normal"},
		{24.99, "Peso normal"},
		{25, "Sobrepeso"},
		{29.99, "Sobrepeso"},
		{30, "Obesidade grau I"},
		{32.65, "Obesidade grau I"},
		{35, "Obesidade grau II"},
		{39.99, "Obesidade grau II"},
		{40, "Obesidade grau III"},
		{55, "Obesidade grau III"},
	}

	for _, c := range casos {
		if got := Classificar(c.imc); got != c.quer {
			t.Errorf("Classificar(%v) = %q, quer %q", c.imc, got, c.quer)
		}
	}
}

func TestFormatar(t *testing.T) {
	casos := []struct {
		imc  float64
		quer string
	}{
		{32.6530612, "32.65"},
		{21.6049382, "21.60"},
		{18.5, "18.50"},
	}

	for _, c := range casos {
		if got := Formatar(c.imc); got != c.quer {
			t.Errorf("Formatar(%v) = %q, quer %q", c.imc, got, c.quer)
		}
	}
}
