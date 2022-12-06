#! /usr/bin/bash

#output the given input
#echo "this is me starting the bash programming language"


# taking the input from the user.
#This method will overwrite the file
#cat>file.txt

#Method to append in the file 
#cat >>file.txt

# single line command

: '
this is the first comment
this is the second comment
this is the third comment
this is the fourth comment
'
# Here doc used for printing the comment in the output(we can choose any variable instead of kreative)
: '
cat << kreative
this is the 1 line comment
this is the second line comment

kreative
'

#If else
: '
count=78
if ((count <10))
then 
	echo "the age is count"
elif ((count >20 && count <30))
then
	echo "the age is not valid"
else
	echo "the age is not in a record"
fi
'

#While Loop
#Syntax 1
: '
count=1
echo "this is the loop running from 1 to 10"
while [ $count -lt 10 ] 
do 
	
	echo "-------------------------------------"
	echo "$count"
	count=$((count+1))
done
'
#Syntax 2
: '
count=1
echo "code running from 1 to 10"
while ((count<=10))
do 
	echo "---------------"
	echo "$count"
	count=$((count+1))
done
'

#Until Loop-----> This will loop runs until our coundition become True
: '
number=1
until ((number >10))
do 
	if ((number%2==0))
	then
		echo $number
	fi
	number=$((number +1))
done
'

#For Loop
#Syntax 1
: '
for i in {1..10..2} #{starting point..endingPoint..increment}
do 
	echo $i
done

'
#Syntax 2
: '
for ((i =1; i<=10; i++))
do 
	if ((i==5))
	then
		continue
	elif ((i==9))
	then
		break
	fi
	echo $i
done
'

# Taking Input from the user
#args=("$@")
#echo ${args[0]} ${args[1]} ${args[2]}# Printing First Three input
: '
echo $# # Pinting len of the args

echo $@ # Print complete args
'

#Read The file using 
#Syntax1
: '
while read line
do 
	echo $line
done <file.txt

'
#Syntax 2
: '
cat file.txt | while read line
do 
	echo $line
done
'
# Array 
arr1=('kali' 'ubuntu' 'linux')
arr1[3]="mac"
echo "${arr1[0]}"
echo "${arr1[1]}"
echo "${arr1[2]}"
echo "${arr1[@]}"
echo "${#arr1[0]}"

unset arr1[0] # Deletion for the variable





















