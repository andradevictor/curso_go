package main

import (
	"errors"
	"fmt"
)

func main() {
	// inteiros
	// int sózinho cria de acordo com a arquitetura do computador (32, 64bits)
	// int, int8, int16, int32, int64

	var numero int64 = 100
	fmt.Println(numero)

	// unsygned int
	var numero2 uint = 100
	fmt.Println(numero2)

	//alias
	//int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//int8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	//float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000.45
	fmt.Println(numeroReal2)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// aspas simples gera o número do registro na tabela Ascii
	char := 'B'
	fmt.Println(char)

	// todo tipo de dado no GO tem o seu valor 0 que é o valor inicial que nao tem implícito
	// string gera o conteúdo em branco
	// numérico gera o valor 0
	// booleano é falso
	// error é nil
	var texto string
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
