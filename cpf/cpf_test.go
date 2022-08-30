package cpf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatar(t *testing.T) {
	cpfValido := CPFDidatico{NumeroCPF: "34859867394"}
	want := "348.598.673-94"
	if got := cpfValido.Formatar("34859867394"); got != want {
		t.Errorf("Formatar() = %q, want %q", got, want)
	}
}

func TestNumeros(t *testing.T) {
	want := [11]int{3, 4, 8, 5, 9, 8, 6, 7, 3, 9, 4}
	if got := numeros("34859867394"); got != want {
		t.Errorf("Numeros() = %q, want %q", got, want)
	}
}

func TestNaoPodeSerTudoZeroSemFormatacao(t *testing.T) {
	cpfInvalido := CPFDidatico{NumeroCPF: "00000000000"}
	want := false
	if got := cpfInvalido.Valido(); got != want {
		t.Errorf("NaoPodeSerTudoZeroSemFormatacao() = %v, want %v", got, want)
	}
}

func TestCPF(t *testing.T) {
	for _, tt := range []struct {
		TestName       string
		Evaluated      CPFDidatico
		ExpectedResult bool
		Message        string
	}{
		{
			TestName:       "Verificando que o CPF é valido",
			Evaluated:      CPFDidatico{"099.285.890-97"},
			ExpectedResult: true,
			Message:        "O CPF deveria ser válido",
		},
	} {
		assert.Equalf(t, tt.Evaluated.Valido(), tt.ExpectedResult, tt.Message)
	}

}
