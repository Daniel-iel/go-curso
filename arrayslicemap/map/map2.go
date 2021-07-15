package main

import "fmt"

func main() {
	funcsESalario := map[string]float64{
		"Daniel": 11235.45,
		"Maria":  14235.45,
		"Pedro":  15235.45,
	}

	funcsESalario["Peter"] = 1350.0
	delete(funcsESalario, "Inexisteste")

	for nome, salario := range funcsESalario {
		fmt.Println(nome, salario)
	}

}
