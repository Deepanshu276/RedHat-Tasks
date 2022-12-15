package RandomNumber

import (
	"fmt"
	TaskOne "goTask2/question1"
	TaskTwo "goTask2/question2"
	TaskThree "goTask2/question3"
	"math/rand"
	"time"
)

func RandomNumber(inp int) {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(100) + 1
	fmt.Printf("random function called and it's value is %v\n ", randomNumber)
	if inp == 1 {
		TaskOne.TaskOne(randomNumber)
	} else if inp == 2 {
		TaskTwo.TaskTwo(randomNumber)
	} else {
		TaskThree.TaskThree(randomNumber)
	}
}
