package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Посчитать количество билетов между счастливыми, у которых в промежутке больше всего обычных")
	fmt.Println("-----------")
	maxCounter := 0
	luckyTicket := 100000
	var maxTicket int
	for ticket := 100000; ticket < 1000000; ticket++ {
		var result1, result2, counter int
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
		}
		if counter > maxCounter {
			maxCounter = counter
			maxTicket = ticket
		}
	}
	fmt.Printf("Максимальное количество билетов между %d и %d составляет - %d\n", maxTicket, maxTicket-maxCounter, maxCounter)
}
