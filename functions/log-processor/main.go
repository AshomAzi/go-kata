package main

import "fmt"

func LogBatch(prefix string, messages ...string) {
	defer fmt.Println("[SYSTEM] Batch processing complete.")
	for _, v := range messages {
		fmt.Printf("[%v] %v\n",prefix, v)
	}
}

func main() {
	LogBatch("Hello", "Jane", "Peter", "Peter")
}
