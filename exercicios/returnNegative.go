package main

import (
	"fmt"
)

func main() {
	fmt.Println(RetornaNegativo(3))
	fmt.Println(RetornaNegativo(0))
	fmt.Println(RetornaNegativo(-4))
}


func RetornaNegativo(num int) int {
	if num <= 0 {
		return num
	}
	return num * (-1)

}
