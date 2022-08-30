package cpf

import (
	"regexp"
	"strconv"
	"strings"
)

type CPF interface {
	Numero() string
	Formatar(numero string) string
	RemoverFormatacao(numero string) string
	Valido() bool
}

type CPFDidatico struct {
	NumeroCPF string
}

var _ CPF = CPFDidatico{} // Assegura a implementação da interface.

func (c CPFDidatico) Numero() string {
	return c.NumeroCPF
}

func (c CPFDidatico) Formatar(numero string) string {
	var numeroSemFormatacao = c.RemoverFormatacao(numero)

	return numeroSemFormatacao[0:3] + "." +
		numeroSemFormatacao[3:6] + "." +
		numeroSemFormatacao[6:9] + "-" +
		numeroSemFormatacao[9:11]
}

func (c CPFDidatico) RemoverFormatacao(numero string) string {
	var regexp = regexp.MustCompile("\\D")
	return regexp.ReplaceAllString(c.NumeroCPF, "")
}

func (c CPFDidatico) Valido() bool {
	if !informouNumero(c.Numero()) {
		return false
	}

	if !estaNoPadrao(c.Numero()) {
		return false
	}

	if numeroInvalido(c.Numero()) {
		return false
	}

	return osDigitosVerificadoresSaoIguais(numeros(c.Numero()))
}

func informouNumero(numero string) bool {
	if len(numero) > 0 {
		return true
	}
	return false
}

func estaNoPadrao(numero string) bool {
	cpfSemMascara, _ := regexp.MatchString("^\\d{11}$", numero)
	cpfComMascara, _ := regexp.MatchString("^\\d{3}.\\d{3}.\\d{3}-\\d{2}$", numero)

	return cpfSemMascara || cpfComMascara
}

func numeroInvalido(numero string) bool {
	numeroSemFormatacao := CPFDidatico{}.RemoverFormatacao(numero)
	for i := 0; i < 10; i++ {
		if numeroSemFormatacao == strings.Repeat(strconv.Itoa(i), 9) {
			return true
		}
	}
	return false
}

func numeros(numero string) [11]int {
	numeroSemFormatacao := CPFDidatico{NumeroCPF: numero}.RemoverFormatacao(numero)
	var numerosString = strings.Split(numeroSemFormatacao, "")
	var numeros [11]int

	for i := 0; i < len(numerosString); i++ {
		numeros[i], _ = strconv.Atoi(numerosString[i])
	}
	return numeros
}

func calcularPrimeiroDigitoVerificador(numeros [11]int) int {

	var resultadoMultiplicacaoCPFSemDigitoVerificador [9]int

	for i := 0; i < 9; i++ {
		resultadoMultiplicacaoCPFSemDigitoVerificador[i] = numeros[i] * (10 - i)
	}

	somaMultiplicacao := 0
	for i := 0; i < 9; i++ {
		somaMultiplicacao = somaMultiplicacao + resultadoMultiplicacaoCPFSemDigitoVerificador[i]
	}

	resto := somaMultiplicacao % 11

	return 11 - resto
}

func calcularSegundoDigitoVerificador(numeros [11]int) int {
	multiplicadoresCPFComUmDigitoVerificador := [10]int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	var resultadoMultiplicacaoCPFComUmDigitoVerificador [10]int

	for i := 0; i < 10; i++ {
		resultadoMultiplicacaoCPFComUmDigitoVerificador[i] = numeros[i] * multiplicadoresCPFComUmDigitoVerificador[i]
	}

	somaMultiplicacao := 0
	for i := 0; i < 10; i++ {
		somaMultiplicacao = somaMultiplicacao + resultadoMultiplicacaoCPFComUmDigitoVerificador[i]
	}

	resto := somaMultiplicacao % 11

	return 11 - resto
}

func osDigitosVerificadoresSaoIguais(numeros [11]int) bool {
	primeiroDigitoVerificador := numeros[9]
	segundoDigitoVerificador := numeros[10]

	return (primeiroDigitoVerificador == calcularPrimeiroDigitoVerificador(numeros)) &&
		(segundoDigitoVerificador == calcularSegundoDigitoVerificador(numeros))
}
