package main

import "fmt"

func main() {

	// idade := map[string]int{}

	// idade["Din Djarin"] = 35
	// idade["Obi-wan kenobi"] = 32

	// fmt.Println(idade)
	// fmt.Println(idade["Din Djarin"])
	// fmt.Println(idade["Obi-wan kenobi"])

	anoNasc := map[string]int{
		"anakin": 1030,
		"yoda": 200,
	}

	fmt.Println(anoNasc)
	fmt.Println(anoNasc["anakin"])
	fmt.Println(anoNasc["yoda"])
	anoNasc["golang-do-zero"] = 2024
	fmt.Println(anoNasc)


}