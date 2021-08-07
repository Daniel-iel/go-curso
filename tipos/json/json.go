package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID   int    `json:"id"`
	Nome string `json:nome`
}

func main() {
	p1 := produto{1, "Notebook"}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 produto
	jsonString := `{"id":1,"Nome":"Notebook"}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Nome)

}
