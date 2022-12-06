#! /usr/bin/bash
#Method 1
#Time Complexity -----------> O(n)

arr=($(cat ques3.txt | sed  's/[^0-9]//g'))
echo "${arr[@]}"
sum=0
tot=0
for i in ${arr[@]}; 
do
  let tot+=$i
done
echo "Total: $tot"

#Method 2
awk 'BEGIN{sum=0} {sum=sum+$1} END {print sum}' ques3.txt
