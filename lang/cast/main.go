package main

import "fmt"

func main() {
	// static cast
	var i int = 50
	var f float64 = float64(50)

	fmt.Println(i, f)

	// string to []byte, and reverse
	var s string = "(´ﾟ◞౪◟ﾟ｀)"
	var arr []byte = []byte(s)
	s = string(arr)

	fmt.Println(s, arr)

	// Type Assesion
	var obj interface{} = "test"

	asStr, succeed := obj.(string)
	fmt.Println(asStr, succeed)

	asInt, succeed := obj.(int)
	fmt.Println(asInt, succeed)
}
