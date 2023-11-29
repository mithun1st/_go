package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

func main() {
	var data string = "Hello World"
	var pipeLine string = "a|b|c"
	var list []string = []string{"x", "y", "z"}

	fmt.Println(reflect.TypeOf(data))
	fmt.Println(utf8.RuneCountInString(data))

	fmt.Println(strings.Split(pipeLine, "|"))
	fmt.Println(strings.Join(list, "#"))
}
