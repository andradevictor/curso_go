package main

import "fmt"

var n int

func init() {
	n = 10
	fmt.Println("Função init sendo executada")
}

func main() {
	fmt.Println("Função main sendo executada", n)
}
