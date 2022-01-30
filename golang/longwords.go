package main

import "fmt"

func main() {
	var times int
	fmt.Scanln(&times)
	words := make([]string, times)
	for i := 0; i < times; i++ {
		fmt.Scanln(&words[i])
	}

	for _, word := range words {
		if len(word) <= 10 {
			fmt.Println(word)
		} else {
			fmt.Printf("%c%d%c\n", word[0], len(word)-2, word[len(word)-1])
		}
	}
}
