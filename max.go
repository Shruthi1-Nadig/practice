package main

import "fmt"
func main () {
var num1 int = 6
var num2 int = 16
result := max(num1,num2) 
fmt.Println("max of num1 and num 2 is ", result)
}

func max(num1,num2 int) int {
	var result int
	if num1 > num2 {
	result = num1
	} else {
		result = num2
	}
return result
}

