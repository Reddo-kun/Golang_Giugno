package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var slice []int
	lista := os.Args[1:]
	for _, numero := range lista {
		n, _ := strconv.Atoi(numero)
		slice = append(slice, n)
	}

	prodottocoppie := smistamento(slice)
	somma := calcola(prodottocoppie)
	fmt.Print("la somma dei prodotti pari è ", somma)

}
func smistamento(sl []int) []int {
	var sliceprodotti []int
	ncoppie := 0
	for i := 0; i < len(sl); i++ {
		for j := i + 1; j < len(sl); j++ {
			ncoppie++
			sliceprodotti = append(sliceprodotti, sl[i]*sl[j])
		}

	}

	return sliceprodotti
}

func calcola(pc []int) int {
	var somma int
	for i := 0; i < len(pc); i++ {
		if pc[i]%2 == 0 {
			somma += pc[i]
		}
	}
	return somma
}
