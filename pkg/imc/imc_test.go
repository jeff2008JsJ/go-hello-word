package imc

import (
	"errors"
	"testing"
)

func TestCalcular(t *testing.T) {
	valor, err := Calcular(100, 1.75)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if got := Formatar(valor); got != "32.65" {
		t.Fatalf("IMC formatado = %q, esperado %q", got, "32.65")
	}
}

func TestCalcularEntradasInvalidas(t *testing.T) {
	if _, err := Calcular(0, 1.75); !errors.Is(err, ErrPesoInvalido) {
		t.Fatalf("peso zero: err = %v, esperado ErrPesoInvalido", err)
	}
	if _, err := Calcular(80, 0); !errors.Is(err, ErrAlturaInvalida) {
		t.Fatalf("altura zero: err = %v, esperado ErrAlturaInvalida", err)
	}
}

func TestClassificar(t *testing.T) {
	casos := map[float64]string{
		17:   "Abaixo do peso",
		22:   "Peso normal",
		27:   "Sobrepeso",
		32:   "Obesidade grau I",
		37:   "Obesidade grau II",
		45:   "Obesidade grau III",
		18.5: "Peso normal",
		25:   "Sobrepeso",
	}
	for valor, esperado := range casos {
		if got := Classificar(valor); got != esperado {
			t.Errorf("Classificar(%v) = %q, esperado %q", valor, got, esperado)
		}
	}
}

func TestResumo(t *testing.T) {
	got, err := Resumo(100, 1.75)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if esperado := "IMC: 32.65 (Obesidade grau I)"; got != esperado {
		t.Fatalf("Resumo = %q, esperado %q", got, esperado)
	}
}
