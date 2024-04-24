package main

import (
	"fmt"
	"time"
)

// func NamedReturn(value int) (result int) {
// 	result = 1

// 	if value == 1 {
// 		return // This has the effect of returning result
// 	}

// 	go func() {
// 		value = result // This has the effect of assigning result to value
// 	}()
// 	return 20 // This has the effect of returning result=20
// }

// func main() {

//		fmt.Println("Value of result", NamedReturn(1))
//		fmt.Println("Value of result", NamedReturn(0))
//		fmt.Println("Value of result", NamedReturn(20))
//	}
var value *int

func NamedReturn() (result int) {
	// if *value == 0 {
	// 	return // This has the effect of returning result
	// }
	
	go func() {
		// result is default value 0
		// value = &result
		result = 10 // This has the effect of assigning result to value
	}()

	time.Sleep(time.Second * 2)
	return 20 // This has the effect of returning result=20
}

func main() {
	
	val := 0
	value = &val
	fmt.Println("Value of result: ", NamedReturn(), " value: ", *value) // This will return 10
	time.Sleep(time.Second)
	val = 1
	value = &val
	fmt.Println("Value of result: ", NamedReturn(), " value: ", *value) // This will return 10
	val = 2
	value = &val
	time.Sleep(time.Second)
	fmt.Println("Value of result: ", NamedReturn(), " value: ", *value) // This will return 10
}
