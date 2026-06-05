package main

import ("fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)




func TempConverter(num float64, from, to string) float64 {
	val := 0.0
	if strings.ToLower(from) == "celcius" && strings.ToLower(to) == "faranhite" {
		val =  (num*2)+30
	}

	if strings.ToLower(from) == "celcius" && strings.ToLower(to) == "kelvin" {
		val =  num+273.15
	}

	if strings.ToLower(from) == "farenhite" && strings.ToLower(to) == "celcius" {
		val =  (num-30)/2
	}

	if strings.ToLower(from) == "farenhite" && strings.ToLower(to) == "kelvin" {
		val =  (num + 459.67) * 5/9
	}

	if strings.ToLower(from) == "kelvin" && strings.ToLower(to) == "celcius" {
		val =  num - 273.15
	}

	if strings.ToLower(from) == "kelvin" && strings.ToLower(to) == "farenhite" {
		val =  (num * 9/5 - 459.67)
	} else {
		return 0.0
	}
	return val
}

func main() {
	fmt.Print("Input the value, what you are converting from kelvin, celcius or farenhite and what you are converting to celcius, kelvin or farenhite respectively: ")
	num := bufio.NewScanner(os.Stdin)
	num.Scan()
	line := num.Text()
	each := strings.Fields(line)
	if len(each) == 3 {
		num, err := strconv.ParseFloat(each[0], 64)
		if err == nil {
			fmt.Println(TempConverter(num, each[1], each[2]))
		} else {
			fmt.Println("Input a valid digit")
		}
	} else {
		fmt.Println("Exactly 3 input is required!")
	}


}