package main 

import "fmt"

func main(){
//map é uma estrutura de dados
	menu := map[string]float64{
//map escolhe o tipo de dado da chave e do valor
		"pizza":  40.00,
		"suco":   12.50,
		"x-tudo": 22.90,
	}
	//fmt.Println(menu["pizza"])

	for k,v := range menu {
	
		fmt.Println("Comida:",k,"R$",v)
	}

	contanova := novaConta("Júlia")
	fmt.Println(contanova)

	contanova2 := novaConta("Milena")
	fmt.Println(contanova2)
}

