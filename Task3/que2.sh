#Method 1 Using sort and uniq (This will Not maintain the order in which the data was stored in orginal file)
#Time Complexity -----------> O(nlogn)

sort que2.txt|uniq

#Method 2 Using awk 

#Time Complexity -----> O(n)
awk '!a[$0]++' que2.txt








