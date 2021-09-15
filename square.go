package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputSize string
	for {

		fmt.Println("\n------------------------------------------------------------------------")
		fmt.Println("Please enter you size between 1 and 20: (type 'exit' to end the program)")
		fmt.Scanln(&inputSize)

		if inputSize == "exit" {
			break
		}

		size, err := strconv.ParseInt(inputSize, 0, 64)

		if err != nil {
			fmt.Println("We accept only number. Please try again... ")
		} else if size > 1 && size < 20 {
			{
				for row := 0; row < int(size); row++ {
					if row == 0 || row == (int(size)-1) {
						fmt.Println(strings.Repeat("* ", int(size)))
					} else {
						fmt.Print("*")
						fmt.Print(strings.Repeat(" ", 2*int(size)-3))
						fmt.Println("*")
					}
				}
			}
		} else {
			fmt.Println("We accept size between 1 and 20")
		}

	}
}
