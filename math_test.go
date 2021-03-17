package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma inválido: Resultado >>> %d. Esperado >>> %d", total, 30)
	}
}

func TestMultiplica(t *testing.T) {
	total := Multiplica(10, 10)

	if total != 100 {
		t.Errorf("Resultado da multiplicação inválido: Resultado >>> %d. Esperado >>> %d", total, 100)
	}
}
