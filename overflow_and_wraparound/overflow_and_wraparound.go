package main

import "fmt"

import "unicode/utf8"

func main() {
	var maxUnit32 uint32 = 4294967295 // but if we add 1 in the declaration it will throw error at compile time which is overflow

	maxUnit32 += 1 // this is wraparound as we add 1 to the value at runtime, it will wrap around to 0
	fmt.Println(maxUnit32)

	rawLiteral := `Hello, this is a single line 
	this is a multiline string`
	
	fmt.Println(utf8.RuneCountInString(rawLiteral))


	fmt.Println(rawLiteral)

	demoArray := [2] string{"hello", "world"}
	fmt.Println(demoArray)

	demoSlices := [] int{-1, 0, 1, 2, 3}
	demoSlices = append(demoSlices, 4)
	fmt.Println(demoSlices)
}
