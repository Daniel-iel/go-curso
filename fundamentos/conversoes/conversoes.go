package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notafinal := int(nota)
	fmt.Println(notafinal)

	fmt.Println("Teste " + string(97))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string para int
	num, _ := strconv.Atoi("123")
	fmt.Println("Teste " + strconv.Itoa(num))
	fmt.Println(num - 122)

	//string para bolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("é verdadeiro")
	}

}
