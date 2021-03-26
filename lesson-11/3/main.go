/*
Напишите программу, которая выведет все части строки, разделенные пробелами, которые можно привести к целому числу.
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var str = "a10 10 20b 20 30c30 30 dd"
	// var str = "    "
	f := func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsSpace(c)
	}
	fmt.Printf("%q", strings.FieldsFunc(str, f))
}
