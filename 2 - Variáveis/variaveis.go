package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1" // explícito
	variavel2 := "Variável 2"           // implícito - (Inferência de Tipos)

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5", "variavel6"

	fmt.Println(variavel5, variavel6)

	const contante1 string = "Constante1"

	fmt.Println(contante1)

	// inversão de valores de variáveis
	variavel5, variavel6 = variavel6, variavel5
}
