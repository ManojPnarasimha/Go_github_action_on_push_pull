package main

import "fmt"

func addition(a, b int) int {
	return a + b
}

func main() {

	fmt.Println("Github actions workflow!")
	result := addition(5, 6)
	fmt.Printf("The result is %v", result)
}
