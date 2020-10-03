package main

import "fmt"

func main() {

	var withdrawal int

	fmt.Println("Банкомат")
	fmt.Println("--------")
	fmt.Println("Внимание: банкомат выдает купюры только номиналов 100 рублей и не больше 100000 рублей")
	fmt.Println("Лимит на снятие не более 100000 рублей")
	fmt.Println("Введите сумму для снятия")
	fmt.Scan(&withdrawal)
	fmt.Println("........")

	if withdrawal > 100000 {
		fmt.Println("Превышен допустимый лимит на снятие, укажите другую сумму")
	} else if withdrawal < 100 {
		fmt.Println("Указана слишком маленькая сумма для снятия")
	} else if withdrawal%100 != 0 {
		fmt.Println("Указанная сумма не может быть выдана, укажите сумму кратную 100 рублей")
	} else {
		var note int
		for note = 0; note < withdrawal/100; note++ {
			fmt.Println("=========")
			fmt.Println("=  100  =")
			fmt.Println("=========")
			fmt.Println("")
		}
		fmt.Printf("Выдано %d купюр номиналом 100 рублей\n", note)
		fmt.Println("Возьмите ваши деньги. Спасибо!")
	}
}
