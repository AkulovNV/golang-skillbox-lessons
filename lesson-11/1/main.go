/*
Напишите программу, которая выведет количество слов начинающихся с большой буквы.
Попробуйте ввести другие строки, например, пустую строку
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"
	// var str = "   "
	var words int
	for len(str) > 0 {
		letter := string(str[0])
		switch {
		case letter == " ":
			str = strings.TrimPrefix(str, letter)
		case letter == ",":
			str = strings.TrimPrefix(str, letter)
		case letter == ".":
			str = strings.TrimPrefix(str, letter)
		default:
			if strings.Compare(letter, strings.ToUpper(letter)) == 0 {
				words += 1
			}
			str = str[1:]
		}
	}
	fmt.Println("Количество слов с заглавной буквы:", words)
}
