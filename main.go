package main

import (
	"github.com/henribas/cpf/pkg/cpf"
	"github.com/henribas/cpf/pkg/server"
)

func main() {
	//jao := cpf.CPFDidatico{NumeroCPF: "123.456.789-11"}
	//teste(jao)
	srv := server.NewServer()
	srv.Run()
}

func teste(c cpf.CPF) {
	//fmt.Println(c.RemoverFormatacao(c.Numero()))
}
