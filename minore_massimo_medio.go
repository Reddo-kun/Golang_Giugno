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
	med := media(min, max)
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
func media(min int, max int) float64 {
	minflt := float64(min)
	maxflt := float64(max)
	var media float64
	media = (maxflt + minflt) / 2
	return media
}
