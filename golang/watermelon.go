package main

import "fmt"

func main() {
	var weight int
	fmt.Scanln(&weight)

	// Since 2 is the smallest even number,
	// we can't divide 2kg of watermelon into two
	// given that the kids are great fans of even numbers.
	// That's why weight <= 2 should be considered "NO"
	if weight <= 2 || weight%2 != 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
