package main

import (
	"fmt"
	"strconv"
)

func main() {

	//str to int
	const str1 string = "100"
	num1, _ := strconv.ParseInt(str1, 2, 32)
	fmt.Println(num1)
	num2, _ := strconv.ParseInt(str1, 10, 64)
	fmt.Println(num2)

	//str to flaot
	const str2 string = "3.1416"
	num3, _ := strconv.ParseFloat(str2, 32)
	fmt.Println(num3)
	num4, _ := strconv.ParseFloat(str2, 64)
	fmt.Println(num4)

	//####################################

	//int to string
	const num_1 int = 1234
	var str_1 string = strconv.Itoa(num_1)
	fmt.Println(str_1)

	//flot to string
	var num_2 float32 = 9.55
	var str_2 string = fmt.Sprintf("%v", num_2)
	fmt.Println(str_2)

	//####################################

	//flot64 to int
	var pi float64 = 3.1415
	var piInt int = int(pi)
	fmt.Println(piInt)

	//flot64 to flot32
	var pi1 float64 = 3.1415
	var piFloat32 float32 = float32(pi1)
	fmt.Println(piFloat32)
}
