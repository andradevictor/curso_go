package main

import "fmt"

// acesso de variáveis fora da funçao
func closure() func() {
	texto := "Dentro da funcao closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da funcao main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
