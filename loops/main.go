package main

import (
	"fmt"
	"time"
)

func main() {

	// arrays
	//var intArr [3]int32
	//or
	// intArr := [3]int32{1, 2, 3}
	// or
	intArr := [...]int32{1, 2, 3}
	intArr[1] = 14
	fmt.Println("Arrays")
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	// iterating the array
	for index, value := range intArr {
		fmt.Printf("Index %v and Value %v\n", index, value)
	}
	fmt.Printf("\n")

	/// slices
	fmt.Println("Slices")
	var intSlices []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %x", len(intSlices), cap(intSlices))
	intSlices = append(intSlices, 7)
	fmt.Printf("\nThe length is %v with capacity %x\n", len(intSlices), cap(intSlices))

	var intSlices2 []int32 = []int32{8, 9}
	/// we can append another list in slices like below
	intSlices = append(intSlices, intSlices2...)
	fmt.Println(intSlices)

	// here make takes three parameters:
	// `[]int32 takes the type of data,`
	// `3` is the max length of slices
	// `10` is the capacity of slices which is optional
	var intSlices3 []int32 = make([]int32, 3, 10)
	fmt.Println(intSlices3)
	fmt.Printf("The length is %v with capacity %x\n\n", len(intSlices3), cap(intSlices3))

	/// maps
	var myMap map[string]uint8 = make(map[string]uint8)
	println("Maps")
	println(myMap)

	var myMap2 = map[string]uint8{"Ram": 48, "shyam": 30}

	fmt.Println(myMap2["Ram"])
	// map always returns some value if it doesnot exist, so it has a boolean check

	var age, ok = myMap2["Ram"]

	if ok {
		fmt.Printf("The age of ram is %v\n", age)
	} else {
		fmt.Printf("Name not found\n")
	}
	// adding new key value pair
	myMap2["hari"] = 20
	fmt.Println(myMap2)

	// removing key value
	delete(myMap2, "hari")
	fmt.Println(myMap2)

	// iterating the map
	for name, age := range myMap2 {
		fmt.Printf("The age of %v is %v\n\n", name, age)
	}

	// while loops is not defined in go instead
	println("While Loops")
	// var i int = 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	//or
	// for {
	// 	if i >= 10{
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }
	//or
	// for i:= 0; i< 10; i++ {
	// 	fmt.Println(i)
	// }

	// Testing pre allocation and dynamic allocaiton iteration time
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))
	
}


func timeLoop(slice []int, n int) time.Duration {
	var now = time.Now()
	for(len(slice) < n){
		slice = append(slice, 1)
	}
	return time.Since(now)
}