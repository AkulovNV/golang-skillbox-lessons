/*
Напишите программу, которая выведет все части строки, разделенные пробелами, которые можно привести к целому числу.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "a10 10 20b 20 30c30 30 dd"
	// var str = "    "
	var stringNumbers string
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			stringNumbers += string(str[i])
		}
		_, err := strconv.ParseInt(string(str[i]), 10, 8)
		if err != nil {
			continue
		} else {
			stringNumbers += string(str[i])
			fmt.Println(i, string(str[i]), "|", stringNumbers)
		}
	}
	fmt.Println(stringNumbers)
}
