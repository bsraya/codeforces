package main

import "fmt"

func main() {
	var max int = 7
	var input string
	fmt.Scanln(&input)
	length := len(input)
	// Consider from i i+1 i+2 i+3 .. to i + max
	// Since i is included, then total numbers -1
	// i + max - 1
	for i := 0; i <= length-max; i++ {
		var counter int = 0
		for j := i; j <= i+max-1; j++ {
			if input[i] == input[j] {
				counter++
			}
		}
		if counter == max {
			fmt.Printf("YES\n")
			return
		}
	}
	fmt.Printf("NO\n")
}
