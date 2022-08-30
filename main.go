package main

import (
	"fmt"

	"github.com/henribas/cpf/cpf"
)

func main() {
	jao := cpf.CPFDidatico{NumeroCPF: "123.456.789-11"}
	teste(jao)
}

func teste(c cpf.CPF) {
	fmt.Println(c.RemoverFormatacao(c.Numero()))
}
