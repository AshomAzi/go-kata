package main

import ("fmt"
	"strings"
)

const (
	vowels = "aeiouAEIOU"
	consonant = "BCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz"
)

func Analyzer(input string) {

	for _, char := range input {
		if strings.Contains(vowels, string(char)) {
			fmt.Printf("%v is a vowel, ", string(char))
		} else if strings.Contains(consonant, string(char)) {
			fmt.Printf("%v is a consonant, " , string(char))
		} else {
			fmt.Printf("%v is a special charater, ", string(char))
		}
		fmt.Printf("%v, %U %v\n", char, char, []byte(string(char)))
	}
}

func main() {
	Analyzer("We🥰lcome")
}