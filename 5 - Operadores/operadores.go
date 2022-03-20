package main

import "fmt"

func main() {
	// Aritmeticos
	// + - / * %

	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 10 * 5
	divisao := 10 / 4
	restoDeDivisao := 10 % 2

	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDeDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 10

	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)

	//Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	// Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// Operadores lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	// Operadores Unários
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	var texto string

	if numero > 5 {
		texto = "Texto > 5"
	} else {
		texto = "Texto < 5"
	}

	fmt.Println(texto)
}
