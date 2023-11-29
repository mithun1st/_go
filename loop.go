package main

import (
	"fmt"
)

func main() {

	//for loop
	for i := 1; i < 5; i++ {
		fmt.Println("for loop ", i)
	}

	//while loop
	var count int = 10
	for count > 0 {
		fmt.Println("while loop ", count)
		count--
	}

	//infinity break
	count = 10
	for {
		fmt.Println("bla bla ", count)
		count++
		if count > 20 {
			break
		}
	}

	//continue
	for count = 1; count <= 15; count++ {
		if count%2 == 0 {
			continue
		}
		fmt.Println(count)
	}

	//range list
	var list []string = []string{"aa", "bb", "cc"}
	for i, e := range list {
		fmt.Println(i, " ", e)
	}

	//range map
	var mp map[string]int = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for i, e := range mp {
		fmt.Println(i, " ", e)
	}
}
