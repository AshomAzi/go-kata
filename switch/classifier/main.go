package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Classifier(char string) {
	val := strings.ToLower(char)
	switch val {
	case "a", "e", "i", "o", "u":
		fmt.Printf("%v is a vowel!\n", char)
	case ".", ",", "!", ";", "?":
		fmt.Printf("%v is a punctuation!\n", char)
	default:
		fmt.Printf("%v is a consonant or other!\n", char)
	}
}

func main() {

	fmt.Print("Input a character: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	Schar := input.Text()
	single := strings.Fields(Schar)
	if len(single) == 1 && len(single[0]) == 1 {
		Classifier(single[0])
	} else {
		fmt.Println("Invalid Input character!")
	}

}
