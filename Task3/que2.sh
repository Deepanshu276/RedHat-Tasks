#Method 1 Using awk 
#Time Complexity -----> O(n)
awk '!a[$0]++' que2.txt

#Method 2 Using sort and uniq (This will Not maintain the order in which the data was stored in orginal file)
#Time Complexity -----------> O(nlogn)
sort que2.txt|uniq

#Method 3 using sort without uniq command
#Time Complexity------------> O(nlogn)
sort -u que2.txt





