package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	slice := leggitesto()
	sliceparole := separaparole(slice)
	mappa := contaripetizioni(sliceparole)
	stampamappa(mappa)
}
func leggitesto() []string {
	var slice []string
	fmt.Println("inserisci delle stringhe: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo := scanner.Text()
		slice = append(slice, testo)

	}
	return slice
}
func separaparole(slice []string) []string {
	var sliceparole []string
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	stringa := strings.Join(slice, " ")
	slice = strings.Fields(stringa)
	for _, parola := range slice {
		word := reg.ReplaceAllString(parola, "")
		sliceparole = append(sliceparole, word)

	}
	return sliceparole
}
func contaripetizioni(slice []string) map[string]int {
	var itemcount int
	mappa := make(map[string]int)
	for _, item := range slice {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := mappa[item]

		if exist {
			mappa[item]++
			// increase counter by 1 if already in the map
		} else {
			mappa[item] = 1
			// else start counting from 1
			itemcount++
		}
	}
	fmt.Println("Parole distinte: ", itemcount)
	return mappa
}
func stampamappa(mappa map[string]int) {
	for i, value := range mappa {
		fmt.Println(i, ":", value)
	}
}
