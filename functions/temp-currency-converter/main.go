package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TempConverter(num float64, from, to string) (float64, error) {
	val := 0.0
	if strings.ToLower(from) == "celcius" && strings.ToLower(to) == "faranhite" {
		val = (num * 2) + 30
	}else if strings.ToLower(from) == "celcius" && strings.ToLower(to) == "kelvin" {
		val = num + 273.15
	}else if strings.ToLower(from) == "farenhite" && strings.ToLower(to) == "celcius" {
		val = (num - 30) / 2
	}else if strings.ToLower(from) == "farenhite" && strings.ToLower(to) == "kelvin" {
		val = (num + 459.67) * 5 / 9
	}else if strings.ToLower(from) == "kelvin" && strings.ToLower(to) == "celcius" {
		val = num - 273.15
	}else if strings.ToLower(from) == "kelvin" && strings.ToLower(to) == "farenhite" {
		val = (num*9/5 - 459.67)
	} else {
		return 0.0, fmt.Errorf("Invalid Input")
	}
	return val, nil
}

func CurrencyConv(amount float64, from, to string) (float64, error) {
	newFrom := strings.ToUpper(from)
	newTo := strings.ToLower(to)
	val := 0.0
	if strings.ToUpper(newFrom) == "USD" && strings.ToUpper(newTo) == "GBP" {
		val = amount * 0.75
	}else if strings.ToUpper(from) == "USD" && strings.ToUpper(newTo) == "EURO" {
		val = amount * 0.86
	}else if strings.ToUpper(newFrom) == "EURO" && strings.ToUpper(newTo) == "USD" {
		val = amount * 1.16
	}else if strings.ToUpper(newFrom) == "EURO" && strings.ToUpper(newTo) == "GBP" {
		val = amount * 0.86
	}else if strings.ToUpper(newFrom) == "GBP" && strings.ToUpper(newTo) == "USD" {
		val = amount * 1.34
	}else if strings.ToUpper(newFrom) == "GBP" && strings.ToUpper(newTo) == "EURO" {
		val = amount * 1.16
	} else {
		return 0.0, fmt.Errorf("Invalid Input")
	}

	return val, nil
}

func main() {

	fmt.Print("Pick from the available conversions \n1 for Currency Converter \n2 for Temperature Converter \n>> ")
	action := bufio.NewScanner(os.Stdin)
	action.Scan()
	new := action.Text()
	each := strings.Fields(new)
	if len(each) == 1 {
		if each[0] == "1" {
			fmt.Print("Input the value, what you are converting from USD, EURO or GBP and \nwhat you are converting to USD, GBP or EURO respectively: ")
			num := bufio.NewScanner(os.Stdin)
			num.Scan()
			line := num.Text()
			each := strings.Fields(line)
			if len(each) == 3 {
				num, err := strconv.ParseFloat(each[0], 64)
				if err == nil {
					val, _ := CurrencyConv(num, each[1], each[2])
					fmt.Println(val)
				} else {
					fmt.Println("Input a valid digit")
				}
			} else {
				fmt.Println("Exactly 3 input is required!")
			}
		}
		if each[0] == "2" {
			fmt.Print("Input the value, what you are converting from kelvin, celcius or farenhite and \nwhat you are converting to celcius, kelvin or farenhite respectively: ")
			num := bufio.NewScanner(os.Stdin)
			num.Scan()
			line := num.Text()
			each := strings.Fields(line)
			if len(each) == 3 {
				num, err := strconv.ParseFloat(each[0], 64)
				if err == nil {
					val, _ := TempConverter(num, each[1], each[2])
					fmt.Println(val)
				} else {
					fmt.Println("Input a valid digit")
				}
			} else {
				fmt.Println("Exactly 3 input is required!")
			}
		}
	} else {
		fmt.Println("Choose between option 1 and 2")
	}
}
