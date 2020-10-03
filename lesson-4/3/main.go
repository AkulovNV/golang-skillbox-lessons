package main

import "fmt"

func main() {

	var withdrawal int

	fmt.Println("Банкомат")
	fmt.Println("Внимание: банкомат выдает купюры только номиналов 100 рублей и не больше 100000 рублей")
	fmt.Println("Лимит на снятие не более 100000 рублей")
	fmt.Println("Введите сумму для снятия")
	fmt.Scan(&withdrawal)

	if withdrawal > 100000 {
		fmt.Println("Превышен допустимый лимит на снятие, укажите другую сумму")
	} else if withdrawal < 100 {
		fmt.Println("Указана слишком маленькая сумма для снятия")
	} else if withdrawal%100 != 0 {
		fmt.Println("Указанная сумма не может быть выдана, укажите сумму кратную 100 рублей")
	} else {
		fmt.Println("Возьмите ваши деньги. Спасибо!")
		for i := 1; i < withdrawal/100; i++ {
			fmt.Println("=========")
			fmt.Println("=  100  =")
			fmt.Println("=========")
			fmt.Println("")
		}
	}
}
