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
	Valido(numero string) bool
}

type CPFDidatico struct {
	NumeroCPF string
}

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

func (c CPFDidatico) Valido(numero string) bool {
	if !informouNumero(numero) {
		return false
	}

	if !estaNoPadrao(numero) {
		return false
	}

	if numeroInvalido(numero) {
		return false
	}

	return osDigitosVerificadoresSaoIguais(numeros(numero))
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

	if numeroSemFormatacao == "00000000000" ||
		numeroSemFormatacao == "11111111111" ||
		numeroSemFormatacao == "22222222222" ||
		numeroSemFormatacao == "33333333333" ||
		numeroSemFormatacao == "44444444444" ||
		numeroSemFormatacao == "55555555555" ||
		numeroSemFormatacao == "66666666666" ||
		numeroSemFormatacao == "77777777777" ||
		numeroSemFormatacao == "88888888888" ||
		numeroSemFormatacao == "99999999999" {

		return (true)
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
	multiplicadoresCPFSemDigitoVerificador := [9]int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	var resultadoMultiplicacaoCPFSemDigitoVerificador [9]int

	for i := 0; i < 9; i++ {
		resultadoMultiplicacaoCPFSemDigitoVerificador[i] = numeros[i] * multiplicadoresCPFSemDigitoVerificador[i]
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
