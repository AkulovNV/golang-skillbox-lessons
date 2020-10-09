package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Количество зеркальных билетов в диапазоне от 100000 до 999999: ")
	var counter int
	luckyTicket := 100000
	for ticket := 100000; ticket < 1000000; ticket++ {
		var result1, result2 int
		for n := 0; n < 3; n++ {
			pos := ticket / int(math.Pow10(n)) % 10
			result1 += pos
		}
		for n := 5; n >= 3; n-- {
			pos := ticket / int(math.Pow10(n)) % 10
			result2 += pos
		}
		if result1 == result2 {
			counter = ticket - luckyTicket
			luckyTicket = ticket
			fmt.Printf("Счастливый билет: %d, количество билетов до него начиная от предыдущего счастливого: %d\n", luckyTicket, counter)
		}
	}
	// fmt.Println(counter)
}
