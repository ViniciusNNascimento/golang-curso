package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(RepeatStr1(6, "I"))
	// fmt.Println(RepeatStr2(60, "Hello"))
	 fmt.Println(RepeatStr3(60, "Hello"))
}

// func RepeatStr1(repetitions int, value string) string {
// 	if repetitions > 0 {

// 		for i := 1; i < repetitions; i++ {
// 			fmt.Print(value)
// 		}
// 		return value
// 	}
// 	return ""
// }

// func RepeatStr2(repetitions int, value string) string {
// 	var repeatStr string
// 	for i := 0; i < repetitions; i++ {
// 		repeatStr += value

// 	}

// 	return repeatStr
// }

func RepeatStr3(repetitions int, value string) string {

	return strings.Repeat(value, repetitions)
}
