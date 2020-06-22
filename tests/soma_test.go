package tests

import "testing"

type SomaTable struct {
	a      int
	b      int
	result int
}

func TestSoma(t *testing.T) {

	data := []SomaTable{
		{1, 2, 3},
		{2, 2, 4},
		{3, 2, 5},
	}

	for _, valor := range data {
		total := Soma(valor.a, valor.b)

		if total != valor.result {
			t.Error("Resultado invalido")
		}
	}
}
