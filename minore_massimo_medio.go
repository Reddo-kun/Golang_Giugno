package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var slice []int
	valori := os.Args[1:]
	for _, numero := range valori {
		n, _ := strconv.Atoi(numero)
		slice = append(slice, n)
	}
	min := minimo(slice)
	max := massimo(slice)
	med := media(slice)
	fmt.Print("Minore: ", min, "\nMassimo: ", max, "\nMedio: ", med)

}
func minimo(sl []int) int {
	var minimo, flag int
	for i := 0; i < len(sl); i++ {
		if flag == 0 {
			minimo = sl[i]
			flag = 1
		}

		if sl[i] < minimo {
			minimo = sl[i]
		}

	}
	return minimo
}
func massimo(sl []int) int {
	var massimo, flag int
	for i := 0; i < len(sl); i++ {
		if flag == 0 {
			massimo = sl[i]
			flag = 1
		}

		if sl[i] > massimo {
			massimo = sl[i]
		}

	}
	return massimo
}
func media(sl []int) float64 {
	var somma int
	var media float64
	for i := 0; i < len(sl); i++ {

		somma = somma + sl[i]

	}
	sommaflt := float64(somma)
	divisore := float64(len(sl))
	media = sommaflt / divisore
	return media
}
