package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// dentro da chave tipo do atributo, fora da chave tipo do valor
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	fmt.Println(usuario["nome"])

	// chave do tipo string e valor é do tipo map
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "engenharia",
			"campus": "1",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")

	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gemeos",
	}

	fmt.Println(usuario2)
}
