package main

import "fmt"

func main() {
	func(texto string) {
		fmt.Println(texto)
	}("parametro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido %s %d", texto, 10)
	}("parametro")

	fmt.Println(retorno)
}
