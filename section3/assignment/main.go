package main

import "fmt"

// Assignment: Create a slice of integer with numbers 0 to 10
// check the integer for even or odd and print the output
func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, "is even")
			// fmt.Printf("%d is even\n", number)
		} else {
			fmt.Println(number, "is odd")
			// fmt.Printf("%d is odd\n", number)
		}
	}
}
