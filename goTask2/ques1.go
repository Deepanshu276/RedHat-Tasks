package TaskOne

import "fmt"

func TaskOne(randomNumber int) {
	if randomNumber > 50 {
		fmt.Print("It's closer to 100\n")

	} else {
		fmt.Print("It's closer to 0\n")
	}
}
