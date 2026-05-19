package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func Grade(input int) {
	switch {
	case input == 100:
		fmt.Println("Pass mark")
		fallthrough
	case input >= 90:
		fmt.Println("Excellent grade!")
	case input >= 80:
		fmt.Println("Excellent")
	default:
		fmt.Println("Fail!!")
	}
}

func main() {
	fmt.Print("Input your score: ")
	snum := bufio.NewScanner(os.Stdin)
	snum.Scan()
	fakenum := strings.Fields(snum.Text())
	if len(fakenum) == 1{
		num, err := strconv.Atoi(fakenum[0])
		if err == nil {
			Grade(num)
		} else {
			fmt.Println("Input a valid digit!")
		}
	} else {
		fmt.Println("Input a valid digit!")
	}
}