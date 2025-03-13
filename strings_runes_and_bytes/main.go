package main

import (
	"fmt"
	"strings"
)

func main() {

	// rune usecase
	var myString = []rune("résumé")
	fmt.Println(len(myString))
	var firstIndex = myString[0]
	fmt.Printf("%v, %T\n", firstIndex, firstIndex)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	// string concatination
	var strSLice = []string{"hello,", "this", "is", "a", "good", "day"}

	// var catStr = ""
	// for i := range strSLice {
	// 	catStr += strSLice[i] + " "
	// }
	// or we can use built in string function which is recommended and builts string after loop is ended
	var strBuilder strings.Builder
	for i := range strSLice {
		strBuilder.WriteString(strSLice[i] + " ")
	}
	fmt.Println(strBuilder.String())

}
