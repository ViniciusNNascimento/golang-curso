package main

import "fmt"

// import "fmt"

func main() {
	// fmt.Println(Summation(8))
	fmt.Println(Summation(2))
}

// func Summation(n int) int {
// 	sum := 0

// 	for i := 1; i <= n; i++ {
// 		sum += i
// 	}
// 	return sum
// }

// outra solução mais simples

func Summation(n int) int {
	return n * (n + 1) / 2
}
