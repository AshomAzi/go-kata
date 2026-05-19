package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Status(input string) {
	switch input {
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("Get ready")
	case "green":
		fmt.Println("Go")
	default:
		fmt.Println("Invalid entery")
	}
}

func main() {
	fmt.Print("Choose a color: ")
	color := bufio.NewScanner(os.Stdin)
	color.Scan()
	fColor := strings.Fields(color.Text())
	if len(fColor) == 1 {
		Status(strings.ToLower(fColor[0]))
	}
}
