package main

import "fmt"

func main() {
	valor := 8

	if valor == 7 {
		fmt.Println("O valoe eh igual a 7")
	} else {
		fmt.Println("O valor eh diferente de 7")
	}

	fmt.Println(1 + 1)
	fmt.Println(5 - 1)
	fmt.Println(2 * 2)
	fmt.Println(8 / 2)
	fmt.Println(8 % 2)

	numero := 8
	if numero%2 == 0 {
		fmt.Printf("%d é par\n", numero)
	} else {

		fmt.Printf("%d é impar\n", numero)
	}

}
