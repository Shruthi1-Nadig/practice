package main

import "fmt"
import "shru/lib"
func main() {
fmt.Println("snack map values before adding a new entry is",lib.Snacks);
lib.Snacks["bread"] = "sour bread"
fmt.Println("snack map values after adding a new entry is",lib.Snacks);
}
