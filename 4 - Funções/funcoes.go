package main

import "fmt"

func main() {
	soma := somar(10, 20)

	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("texto da função1")
	fmt.Println(resultado)

	resultadosSoma, resultadoSubtracao := calculosMatematicos(10, 15)

	// quando passa o _ ele ignora o segundo resultado da função!!
	// ordem dos fatores altera o produto
	// resultadosSoma, _ := calculosMatematicos(10, 15)

	fmt.Println(resultadosSoma, resultadoSubtracao)
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// 2 retornos
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}
