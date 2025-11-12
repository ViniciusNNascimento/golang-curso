package main

import (
	"fmt"
)

// array - tamanho fixo , slice - sem tamanho fixo
// ambos sao tipo unicos

// map = pode misturar tipos

func main() {

	// var array [2]string
	// array[0] = "hello"
	// array[1] = "brother"

	// fmt.Println(array[0])
	// fmt.Println(array[1])
	// fmt.Println(array)

	// numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(numPrimos)
	// fmt.Println(numPrimos[0:4])
	// fmt.Println(numPrimos[1:])

	//var slice []string

	slice := make([]string, 5)

	slice[0] = "Yoda"
	slice[1] = "Mandalorian"

	// fmt.Println(slice[0], slice[1])
	// fmt.Println(slice[0])
	// fmt.Println(slice[1])
	// fmt.Println(slice[2])
	// slice[2] = "vinic"
	// fmt.Println(slice[2])
	// fmt.Println(slice)

	// fmt.Println(len(slice))
	// fmt.Println(slice[2])
	// fmt.Println(slice[3])
	// fmt.Println(slice[4])
	// fmt.Println(slice[5]) exemplo de erro

	numPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)

	numPares = append(numPares, 14, 16, 18, 20)
	fmt.Println(numPares)

}
