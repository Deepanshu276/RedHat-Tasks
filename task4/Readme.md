## Go Task1

```
1. Write a loop to ask the name , age and nationality from a user and print "Your name is ___  your age is ___ years old and your nationality is __"
    i. This should be infinite loop

2. Define the following array "Menu"
    i. Append to it the following item: "hamburger"
    ii. Append to it the following item: "salad"
    iii. Iterate over the list and print for each item Food: <Food name>. Make sure to replace <Food name> with item from the array

3. Define an array of 5 items
    i.Iterate over it and print for each item the following: This is <ITEM> and its index in the array is <INDEX>
    
```
- How to run : `go run filename.go`

## Solution 1
- Time Complexity : 
```
Time Complexity = Undefined
         As f(n)=O(g(n))
         f(n)==infinity , thus it will not map n to real number,hence f(n) is not a real valued function
```
- Space Complxity :
```
Space Complexity = O(1) (The variable declared go out of scope everytime, they get destroyed , hence space complexity remains constant)
```

## Solution 2
- Time Complexity : `O(n)`
- Space Complxity : `O(n)`
- Code Explanation : 
```
1. Created Slice: menu := '[]string{"pasta", "Macaroni", "pizza "}'
2. Appended the slice := 'menu = append(menu, "hamburger", "salad")'
3. Loop through the slice to print the element
```

## Solution 3
- Time Complexity : `Time Complexity = O(n)`(Loop traversing the size of array of n)
- Space Complexity : `Time Complexity= O(n)` (Array of size n is required to store the elements)
