package TaskThree

import "fmt"

func TaskThree(randomNumber int) {
	if randomNumber > 50 && randomNumber%2 == 0 {
		fmt.Print("It's closer to 100 and it's even\n")
	} else if randomNumber < 50 {
		fmt.Println("Random number is not greater than 50 , try again")
	} else {
		fmt.Println("Random number is not even, Try again")
	}
}
