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
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var n int
	var leftBkt = "("
	var rightBkt = ")"
	var b bytes.Buffer
	var cur string

	fmt.Printf("Введите число вариаций скобок: ")
	fmt.Scan(&n)
	// var curBkt = leftBkt
	for j := 1; j <= n; j++ {
		cur = strings.Repeat(leftBkt, j) + strings.Repeat(rightBkt, j)
		switch {
		case len(cur) < n*2:

		}
		for len(cur) < n*2 {
			cur = string(b.Bytes())
			fmt.Println(cur, len(cur))

			switch {
			// case strings.Count(cur, rightBkt)+strings.Count(cur, leftBkt) == n*2:
			// 	return
			case strings.Count(cur, rightBkt) == strings.Count(cur, leftBkt):
				b.WriteString(leftBkt)
			case strings.Count(cur, leftBkt) == n:
				b.WriteString(rightBkt)
			case strings.Count(cur, leftBkt) == j:
				b.WriteString(rightBkt)
			case strings.Count(cur, rightBkt) == j:
				b.WriteString(leftBkt)
			}

		}
		fmt.Printf("%s\n", cur)
		b.Reset()
		// stringBkt = cur
		// curBkt = leftBkt
	}
	// strBktS = strings.Split(strBkt, curBkt)
}
