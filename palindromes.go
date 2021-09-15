package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var inputPalindrome string
	var isPalindrome bool
	var inputLen int

	for {

		fmt.Println("\n-------------------------------------------------------------")
		fmt.Println("Please enter you palindrome: (type 'exit' to end the program) ")
		fmt.Scanln(&inputPalindrome)

		if inputPalindrome == "exit" {
			break
		}

		inputLen = utf8.RuneCountInString(inputPalindrome)
		for i := 0; i < inputLen; i++ {
			if string(inputPalindrome[i]) == string(inputPalindrome[inputLen-(i+1)]) {
				isPalindrome = true
				continue
			} else {
				isPalindrome = false
				fmt.Println("Your in put is not palindrome")
				break
			}
		}

		if isPalindrome {
			fmt.Println("Your input is palindrome")
		}
	}
}
