package main

//Entry Point of the Program and tells the Go compiler that the package should compile as an executable program instead of a shared library.
import "fmt"

func main() {
	var day = [5]string{"Sunday", "Tuesday", "Thursday", "Saturday", " wednesday"}
	/*      Method 1
		Time Complexity ----------> O(n)----Loop traversing the size of array of n
		Space Complexity----------->O(n)----Array of size n is required to store the elements
	*/
	for i := range day {
		fmt.Printf("This is %v and its index in the array is %v\n", day[i], i)
	}

	/*
		Method 2
		Time Complexity ----------> O(n)----Loop traversing the size of array of n
		Space Complexity----------->O(n)----Array of size n is required to store the elements
	*/
	for index, d := range day {
		fmt.Printf("This is %v and its index in the array is %v\n", d, index)
	}

}
