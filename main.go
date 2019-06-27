package main

import "fmt"
import "./lib"
func main() {
fmt.Println("snack map values before adding a new entry is",lib.snacks);
snacks["bread"] = "sour bread"
fmt.Println("snack map values after adding a new entry is",lib.snacks);
}
