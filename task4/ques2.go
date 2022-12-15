package main
//Entry Point of the Program and tells the Go compiler that the package should compile as an executable program instead of a shared library.

import "fmt"

func main() {

	menu := []string{"pasta", "Macaroni", "pizza "}

	fmt.Printf("The menu for the Party  is  %v\n:", menu)

	menu = append(menu, "hamburger", "salad")

	/*
			Method 1:
			Time Complexity: O(n)-----> looping throgh n element
		    Space Complexity : O(n)-----> N element stored in an array
	*/

	for i := range menu {
		fmt.Printf("Food : %v\n", menu[i])
	}
	/*
			Method 2:
			Time Complexity: O(n)-----> looping throgh n element
		    Space Complexity : O(n)-----> N element stored in an array
	*/
	for _, food := range menu {
		fmt.Printf("Food : %v\n", food)

	}
	/*
			Method 3:
			Time Complexity: O(n)-----> looping throgh n element
		    Space Complexity : O(n)-----> N element stored in an array
	*/
	for i := 0; i < len(menu); i++ {
		fmt.Printf("Food : %v\n", menu[i])
	}

}
