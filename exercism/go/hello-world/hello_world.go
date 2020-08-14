package greeting

import "fmt"

func HelloWorld() string {
	// function that just returns a string
	return "Hello, World!"
}

func main() {
	mystring := HelloWorld()
	// var containing the result of fhe function HelloWorld()
	fmt.Println(mystring)
	// print the contents of the mystring variable

}
