/*
Написать программу, которая на вход принимала бы интовое число и для него генерировала бы все возможные комбинации круглых скобок
Пример:
на вход приходит число 3
на выходе:
["((()))","(()())","(())()","()(())","()()()"]
["(((())))","((()))()","(())(())","(())()()",]
["()()","(())"]
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var leftBkt = "("
	var rightBkt = ")"
	// var b bytes.Buffer
	var str string
	// var strBkt []string

	fmt.Printf("Введите число вариаций скобок: ")
	fmt.Scan(&n)
	for j := n; j > 1; j-- {
		for i := j; i >= 1; i++ {
			str = strings.Repeat(leftBkt, j) + strings.Repeat(rightBkt, i)
			bktCount := strings.Count(str, leftBkt) + strings.Count(str, rightBkt)
			if bktCount == n*2 {
				fmt.Println(str)
				return
			} else if bktCount < n*2 && j == i {
				str += (n*2 - bktCount)
			}
		}
		str = strings.Repeat(leftBkt, j) + strings.Repeat(rightBkt, i)
		if len(str) == n*2 {
			fmt.Println(str)
		} else {
			for i := n - j; i > 0; i++ {
				switch {
				case i*2+len(str)+i*2 == n*2:
					fmt.Println(strings.Repeat(leftBkt+rightBkt, i) + str + strings.Repeat(leftBkt+rightBkt, i))
				case (i*2+len(str) == n*2):
					fmt.Println(strings.Repeat(leftBkt+rightBkt, i) + str)
					fmt.Println(str + strings.Repeat(leftBkt+rightBkt, i))
				case i*2+len(str)+i*2 < n*2:
				}
			}
		}
	}
}
