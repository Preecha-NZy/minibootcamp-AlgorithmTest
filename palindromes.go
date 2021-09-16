package main

import (
	"fmt"
	"strings"
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
		sliceInput := strings.Split(inputPalindrome, "")
		inputLen = len(sliceInput)

		for i := 0; i < inputLen; i++ {
			if string(sliceInput[i]) == string(sliceInput[inputLen-(i+1)]) {
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
