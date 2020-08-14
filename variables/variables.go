package main

import (
	"fmt"
	"reflect"
)

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	fmt.Println(reflect.TypeOf(c).String())
	fmt.Println(reflect.TypeOf(python).String())
	fmt.Println(reflect.TypeOf(java).String())
}
