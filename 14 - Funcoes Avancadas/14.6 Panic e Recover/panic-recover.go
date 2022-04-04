package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Funcao recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("Média é exatamente 6")
}

// defer adia a execução até o ultimo momento possível
func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("pós execucao")
}
