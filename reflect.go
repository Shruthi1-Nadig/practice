package main
import (
	"fmt"
	"reflect"
	)

func main() {
var i int32 = 213453
fmt.Println("type of i", reflect.TypeOf(i))
fmt.Println("value of i", reflect.ValueOf(i))
fmt.Println("value of i", reflect.ValueOf(i).Type())
fmt.Println(" of i", reflect.ValueOf(i).Kind())
}
