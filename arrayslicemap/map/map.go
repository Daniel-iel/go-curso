package main

import "fmt"

func main() {

	aprovados := make(map[int]string)

	aprovados[12345678978] = "maria"
	aprovados[98745612300] = "pedro"
	aprovados[98745612301] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[98745612301])
	delete(aprovados, 98745612300)
	fmt.Println(aprovados[98745612300])
}
