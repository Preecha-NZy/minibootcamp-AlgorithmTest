package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	for {
		var values []int64
		var isDigit bool

		fmt.Println("\n-------------------------------------------------------------")
		fmt.Println("Please enter you number : (type 'exit' to end the program) ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := strings.Split(scanner.Text(), " ")

		if scanner.Text() == "exit" {
			break
		}

		for i := range userInput {
			value, err := strconv.ParseInt(userInput[i], 0, 64)
			if err != nil {
				isDigit = false
				fmt.Println("We accept only number. Please try again... ")
				break
			} else {
				isDigit = true
				values = append(values, value)
			}
		}

		if isDigit {
			min := findMin(values)
			fmt.Println("Your smallest number is:", min)
		}
	}
}

func findMin(values []int64) int64 {
	Min := values[0]
	for i := range values {
		if values[i] < Min {
			Min = values[i]
		}
	}
	return Min
}
