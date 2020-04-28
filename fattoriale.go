package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numero := os.Args[1]
	var slice []int
	n, _ := strconv.Atoi(numero)
	for i := 1; i <= n; i++ {
		fatt := fattoriale(i)
		slice = append(slice, fatt)
	}
	fmt.Println(slice)
}
func fattoriale(n int) int {
	fat := 1
	for i := n; i > 0; i-- {

		fat = fat * i

	}
	return fat
}
