package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var slice []int
	var flag int
	lista := os.Args[1:]
	for _, numero := range lista {
		n, _ := strconv.Atoi(numero)
		slice = append(slice, n)
		flag++
	}
	if flag%2 == 0 {
		calcola(slice)
	} else {
		fmt.Println("errore: numeri inseriti dispari")
	}

}
func calcola(sl []int) {
	var molt int
	for i := 0; i < len(sl); i += 2 {
		molt += sl[i] * sl[i+1]

	}
	fmt.Println("la somma e'", molt)
}
