package main

import "fmt"

func main() {

	// posicao := 2

	// switch posicao {
	// case 1:
	// 	fmt.Println("Primeiro lugar")
	// case 2:
	// 	fmt.Println("Segundo lugar")
	// case 3:
	// 	fmt.Println("Terceiro lugar")

	// }

	nome := "P3PO"
	switch nome {
	case "ayanokoji":
		fmt.Println("É um Personagem")
	case "vinic":
		fmt.Println("É uma Pessoa")
	case "bob":
		fmt.Println("É um Cachorro")
	default:
		fmt.Println("É um alien")
	}
}
