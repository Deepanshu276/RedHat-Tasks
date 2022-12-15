## Go Task2 

Problem Statement

```
1. Generate a random number between 1 and 100
    If the number is higher than 50 print "It's closer to 100"
    If the number is lower than 50 print "It's closer to 0"
    Print the generated random number

2. Modify the previous code to print "It's 50!" if the random number is 50

3. Modify the conditional in the code you previously wrote to check not only if the  number is higher than 50 but also if it's even. If it's even and higher than 50, print,"It's closer to 100 and it's even!"
```

- How to run : `go run filename.go`

- File Structure: 
    - Main.go : 
        ``` Take the input from the user and pass it to random file to generate the random number```
    - Random.go:
        ``` Generate the random number and pass the arguement to file ques1, ques2, ques3 to perform task1 ,task2, task2```
    - ques1.go:
        ``` Perform the Task1  ```
    - ques2.go:
        ``` Perform the Task2 ```
    - ques3.go:
        ```Perform the Task3 ```
