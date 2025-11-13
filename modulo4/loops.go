package main

import (
	"fmt"
)

func main() {

	// sum := 0

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Sum: ", sum)
	// 	fmt.Println("Index: ", i)
	// 	sum += i
	// }

	// fmt.Println(sum)

	// for {
	// 	fmt.Println("Golang")
	// 	time.Sleep(2 * time.Second)
	// }

	frutas := []string{"laraja", "tomate", "maça", "manga"}
	for _, fruta := range frutas {
		fmt.Println("Fruta", fruta)
	}



}
