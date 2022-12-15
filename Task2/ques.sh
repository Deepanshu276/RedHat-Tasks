#! /usr/bin/bash

#Method 1-----> This method will not count spaces although they will count the alphabet present in the file
#Time Complexity: O(N)
#Space Complexity: O(N)

awk '{a[$0]++; if(z<a[$0]){ z=a[$0];s[z]=$0}} END{print s[z]}' IP_addresses.txt

#Method 2----------> This method will not count spaces although they will count the alphabet present in the file
: '
Size of Data: N
Size of Main Memory : M
Merges at a Time: R
Number of Runs created by Sort : N/M
Number of passes through the data is : log(N/M)/log(R))

Time Complexity : Ω((N/M)log(N/M)/log(R))
Space Complxity : O(M)
'
sort IP_addresses.txt | uniq -c | sort -n -r | head -1

#Method 3 ----> This method will not also Count spaces along with alphabet 
: '
-c: It is used to count how many times a line repeated in the searching data and the values are shown as the prefix of that line.
sort : will sort the file in order numbers -> string
sort -nr: Will sort the output in descending order of occurance where count is added as a prefix in to a line
Time Complexity : Ω((N/M)log(N/M)/log(R)) + Ω((N/M)log(N/M)/log(R)) === Ω((N/M)log(N/M)/log(R))
Space Complexity : O(M)
'

sed -e 's/\s/\n/g' < IP_addresses.txt | sort | uniq -c | sort -nr |head -2
