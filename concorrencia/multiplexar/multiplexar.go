package main

import (
	"fmt"

	"github.com/Daniel-iel/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://google.com", "https://youtube.com"),
		html.Titulo("https://amazon.com", "https://microsoft.com"),
	)

	fmt.Println(<-c, "|", <-c)
}
