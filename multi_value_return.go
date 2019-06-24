package main

import "fmt"

func main() {
var x,y = 1,2
swap_x,swap_y := swap(x,y)
fmt.Println("swaped values of x,y", swap_x,swap_y)
}
func swap(x,y int) (int,int) {
return y,x
}

