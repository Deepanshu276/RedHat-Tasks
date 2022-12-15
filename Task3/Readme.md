## Task 3

Problem Statement

1. Write a script to separate day, month, year, hour, minute and second values of the current system date and time.
2. Write a script to remove duplicated lines from a file.
3. Write a script to find the sum of all numbers in a file in Linux

### Solution 1

- How to run : `Bash filename.sh`
- Time Complexity : `O(1)`
- Space Complexity: `O(1)`
```
Explanation: Seperate current Date into Year,Month,Day,Hour,Minute,Second
```

### Solution2 

- How to run : `Bash filename.sh`

- Method 1
    - using sort + uniq
    - Time Complexity : `O(nlogn)`
    - Space Complexity: `O(1)`
    - `Note : This will Not maintain the order in which the data was stored in orginal file`    
- Method 2
    - Using awk
    - Time Complexity : O(n)
    - `Note: This will maintain the Order in which Data was stored in original file`

### Solution3
- How To run: `Bash filename.sh`
- Method 1
    - using awk
    - Time Complexity : O(n)
    - Explanation : `awk 'BEGIN{sum=0} {sum=sum+$1} END {print sum}'`Sums the numbers in the individual lines and finally prints the sum
- Method 2
    - Using array + sed
    - Time Complexity : O(n)
    - Space Complexity: O(n)
    - 
      ```Explanation : Fetched all the number from the file using : sed  's/[^0-9]//g') and stored in the array.Traversed Through the array to find the sum of all the element```
