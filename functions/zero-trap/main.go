package main

import "fmt"

func SafeDivide(numerator, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0.0, fmt.Errorf("Division by zero is no allowed")
	}
	return numerator/denominator, nil
}
func main() {
	result, err := SafeDivide(2,0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
