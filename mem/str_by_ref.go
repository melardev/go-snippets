package mem

import "fmt"

// This snippets shows that strings are passed by value
func printString(myStringPtr string) {
	fmt.Printf("%p\n", &myStringPtr)
}

func printStringPtr(myStringPtr *string) {
	fmt.Printf("%p\n", myStringPtr)
}

func StringByRef() {
	firstStr := "First string"
	fmt.Printf("%p\n", &firstStr)
	printStringPtr(&firstStr)

	secondStr := new(string)
	*secondStr = "Second string"
	fmt.Printf("%p\n", secondStr)
	printStringPtr(secondStr)

}
