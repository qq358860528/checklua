require("table")

print(table.getn{10,2,4})          --> 3

print(table.getn{10,2,nil})        --> 2

print(table.getn{10,2,nil; n=3})   --> 3

print(table.getn{n=1000})          --> 1000



 
a = {n=200}

print(#a)               --> 10

}}
