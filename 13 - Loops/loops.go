package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		//	time.Sleep(time.Second)
		i++
		fmt.Println("Incrementado - i", i)
	}

	for j := 0; j < 10; j++ {
		//time.Sleep(time.Second)
		fmt.Println("Incrementado - j", j)
	}

	nomes := [3]string{"Joao", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println("Indice: ", indice, "Nome: ", nome)
	}

	for _, nome := range nomes {
		fmt.Println("Nome: ", nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "victor",
		"sobrenome": "andrade",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
