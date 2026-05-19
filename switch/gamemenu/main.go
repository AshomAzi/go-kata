package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Menu(input int) {

}

func main() {
outer:
	for {
		fmt.Println("1. Start Game\n2. Options\n3. Exit")
		fmt.Print("Choose an option from the menu: ")
		snum := bufio.NewScanner(os.Stdin)
		snum.Scan()
		num := strings.Fields(snum.Text())
		if len(num) == 1 && len(num[0]) == 1 {
			switch num[0] {
			case "1":
				fmt.Println("Starting Game")
			case "2":
				fmt.Println("Loading Menu")
			case "3":
				break outer
			}
		}
	}
}
