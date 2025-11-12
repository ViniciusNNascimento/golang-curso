package main

import "fmt"

func main() {
	somaDosValores := soma(44, 44)
	fmt.Println("soma: ", somaDosValores)

	fmt.Println("subtração: ", subtracao(44, 6))

	nome1, nome2, _ := printName("vinic")
	fmt.Println(nome1)
	fmt.Println(nome2)
	// fmt.Println(nome3)

	nome, sobrenome := printFullName("anakin", "skywalker")
	fmt.Println(nome, sobrenome)
}


// funcao com a primeira letra minuscula eh privada
func printFullName(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

// funcao com a primeira letra maiuscula eh publica
func PrintName(nome string) string {
	return nome
}

func printName(name string) (string, string, string) {
	return name, name, name
}

func subtracao(x int, y int) int {
	return x - y
}

func soma(x int, y int) int {
	return x + y
}
