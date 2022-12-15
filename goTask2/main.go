package main

import (
	"bufio"
	"fmt"
	RandomNumber "goTask2/randomCode"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Unix is used to yield time as a Unix time that is the number of seconds passed
	userInput := bufio.NewReader(os.Stdin)

	input, _ := userInput.ReadString('\n')
	numRating, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Println("Enter the task number you want to perform!!!!!", numRating)

	switch numRating {
	case 1:
		RandomNumber.RandomNumber(int(numRating))
	case 2:
		RandomNumber.RandomNumber(int(numRating))
	case 3:
		RandomNumber.RandomNumber(int(numRating))
	default:
		fmt.Print("No function called")

	}
}
