package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	slicen := LeggiNumeri()
	ins, suff := FiltraVoti(slicen)
	fmt.Println("insufficienti: ", ins, "\nsufficienti: ", suff)

}
func LeggiNumeri() []int {
	var sliceint []int
	numeri := os.Args[1:]
	for _, n := range numeri {
		num, _ := strconv.Atoi(n)
		sliceint = append(sliceint, num)
	}
	return sliceint
}
func FiltraVoti(slicenumeri []int) ([]int, []int) {
	var insufficienti []int
	var sufficienti []int
	for i := 0; i < len(slicenumeri); i++ {
		if slicenumeri[i] < 60 {
			insufficienti = append(insufficienti, slicenumeri[i])
		} else {
			sufficienti = append(sufficienti, slicenumeri[i])
		}
	}
	return insufficienti, sufficienti
}
