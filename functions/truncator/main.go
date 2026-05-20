package main

import "fmt"

func SecureTruncate(input string, limit int) (string, error) {

	newInput := []rune(input)
	if limit < 1 {
		return "", fmt.Errorf("Limit parameter is out of range.")
	}
	if len(input) > limit {
		return string(newInput[0:limit]) + "...", nil
	}

	return input, nil
}

func main() {
	fmt.Println(SecureTruncate("Ashom", 0))
}
