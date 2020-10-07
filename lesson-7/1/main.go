package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Количество зеркальных билетов в диапазоне от 100000 до 999999: ")
	var counter int
	for ticket := 100000; ticket < 1000000; ticket++ {
		var result1, result2 string
		for n := 0; n < 6; n++ {
			pos := ticket / int(math.Pow10(n)) % 10
			result1 += string(rune(pos))
		}
		for n := 5; n >= 0; n-- {
			pos := ticket / int(math.Pow10(n)) % 10
			result2 += string(rune(pos))
		}
		if result1 == result2 {
			// fmt.Println(ticket)
			counter++
		}
	}
	fmt.Println(counter)
}
