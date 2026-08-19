package winui

import "testing"

func TestAspas(t *testing.T) {
	casos := map[string]string{
		"Peso":            "'Peso'",
		"Digite o 'peso'": "'Digite o ''peso'''",
	}
	for entrada, esperado := range casos {
		if got := aspas(entrada); got != esperado {
			t.Errorf("aspas(%q) = %q, esperado %q", entrada, got, esperado)
		}
	}
}
