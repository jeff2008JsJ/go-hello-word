package main

import "testing"

func TestSaudacao(t *testing.T) {
	if got := saudacao(); got != "Hello, World!" {
		t.Errorf("saudacao() = %q, quer %q", got, "Hello, World!")
	}
}
