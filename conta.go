package main

import "fmt"

type conta struct {
	//conta na vida real
	nome    string
	itens   map[string]float64
	gorjeta float64
}

// retorna a estrutura da comanda acima
func novaConta(nome string) conta {
	//dados da conta
	//criamos variavel c = conta
	c := conta{
		nome:    nome,
		itens:   map[string]float64{},
		gorjeta: 0.0,
	}
	return c
}

func (c conta) formatacao() string {
	fs := "Detalhe da comanda \n"
	var total float64 = 0

	for k, v := range c.itens {
		fs += fmt.Sprintf("%v...R$  %0.2f\n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%v....%0.2f", "total", total)
	return fs
}