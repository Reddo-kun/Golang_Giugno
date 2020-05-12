package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	slice := leggitesto()
	mappa := istogramma(slice)
	letturamap(mappa)
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
func istogramma(slice []string) map[string]int {
	mappa := make(map[string]int)
	stringa := strings.Join(slice, "")
	for _, letter := range stringa {
		if unicode.IsLetter(letter) {
			res1 := strings.Count(stringa, string(letter))
			mappa[string(letter)] = res1
		}

	}
	return mappa
}
func letturamap(mappa map[string]int) {
	fmt.Println("Istogramma: ")
	for i, value := range mappa {
		fmt.Println(i, strings.Repeat("*", value))
	}
}
