/*Time Complexity ---------> Undefined
As f(n)=O(g(n))
f(n)==infinity , thus it will not map n to real number,hence f(n) is not a real valued function

*/
// Space Complexity----------> O(1) (The variable declared go out of scope everytime, they get destroyed , hence space complexity remains constant)
package main

//Entry Point of the Program and tells the Go compiler that the package should compile as an executable program instead of a shared library.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("welcome to the code")
	NAME := bufio.NewReader(os.Stdin)    //Taking name from the User
	AGE := bufio.NewReader(os.Stdin)     //Taking age from the User
	COUNTRY := bufio.NewReader(os.Stdin) // Taking Country of the User

	name, _ := NAME.ReadString('\n')       //Reading the name given by the User till the new line . '\n ' acts as a delimiter
	country, _ := COUNTRY.ReadString('\n') //Reading the country given by the User till the new line . '\n ' acts as a delimiter

	input, _ := AGE.ReadString('\n')

	age, _ := strconv.ParseFloat(strings.TrimSpace(input), 64) // Converting the type of age from string ti float

	for {
		fmt.Printf("Your name is %v Your age is %v Years old and Your nationality is  %v", name, age, country)
	}

}
