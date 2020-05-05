package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var testo string
	var stringa []string
	fmt.Println("inserisci delle stringhe: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo = scanner.Text()
		flag := 0
		for _, char := range testo {
			if unicode.IsDigit(char) && flag == 0 {
				stringa = append(stringa, testo)
				stringa = append(stringa, "\n")
				flag++

			}

		}

	}
	fmt.Print("le stringhe che contengono cifre sono: ")
	flag := 0
	textstring := strings.Join(stringa, "")
	for _, text := range textstring {
		if text == 10 && flag == 0 {
			fmt.Print("\n", textstring)
			flag++
		}

	}

}
