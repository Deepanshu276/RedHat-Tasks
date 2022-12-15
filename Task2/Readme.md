## Task2
Problem Statement

```
you will be given a file with a list of IP addresses called ip_list. 
Using the file, determine which IP address is the most recurring (listed the most times)

```
### Solution 1
Using Hashmap (Bash awk)

How to run : `Bash fileName.sh`

``` 
Time Complexity: O(N)
Space Complexity: O(N)
Note : This method will not count spaces although they will count the alphabet present in the file
```

### Solution 2
Using Bash sort + uniq

How to run : `Bash fileName.sh`

```
Size of Data: N
Size of Main Memory : M
Merges at a Time: R
Number of Runs created by Sort : N/M
Number of passes through the data is : log(N/M)/log(R))
Time Complexity : Ω((N/M)log(N/M)/log(R))
Space Complxity : O(M)
Note : This method will not count spaces although they will count the alphabet present in the file
```
### Solution3 
Using Bash sed+sort+uniq

How to run : `Bash fileName.sh`

```
-c: It is used to count how many times a line repeated in the searching data and the values are shown as the prefix of that line.
sort : will sort the file in order numbers -> string
sort -nr: Will sort the output in descending order of occurance where count is added as a prefix in to a line
Time Complexity : Ω((N/M)log(N/M)/log(R)) + Ω((N/M)log(N/M)/log(R)) === Ω((N/M)log(N/M)/log(R))
Space Complexity : O(M)

```





