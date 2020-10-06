package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	var number = rand.Intn(9) + 1
	var success, more, less = "да", "больше", "меньше"
	var answer string

	fmt.Println("Угадывание числа")
	fmt.Println("Загадате число от 1 до 10 включительно...")
	fmt.Printf("Если это заданное число, напишите без кавычек %q, в противном случае напишите %q или %q ваше число\n", success, more, less)
	time.Sleep(1 * time.Second)

	fmt.Println("Теперь я попробую угадать ваше число")
	fmt.Printf("Вы загадали число: %v\n Так ли это?\n", number)

	fmt.Scan(&answer)

	if answer == success {
		fmt.Println("Ура! :)")
	} else {
		if answer == more && number <= 3 {
			number += 3
			fmt.Printf("Вы загадали число: %v\n Так ли это?\n", number)
			fmt.Scan(&answer)
		} else if answer == more && number > 3 && number < 7 {
			number += 2
			fmt.Printf("Вы загадали число: %v\n Так ли это?\n", number)
			fmt.Scan(&answer)
		} else if answer == more {
			number += 1
			fmt.Printf("Вы загадали число: %v\n Так ли это?\n", number)
			fmt.Scan(&answer)
		} else if answer == less && number <= 3 {
			number -= 1
			fmt.Printf("Вы загадали число: %v\n Так ли это?\n", number)
			fmt.Scan(&answer)
		} else if answer == less && number > 3 && number < 7 {
			number -= 2
			fmt.Printf("Вы загадали число: %v\n Так ли это?\n", number)
			fmt.Scan(&answer)
		} else {
			number -= 3
			fmt.Printf("Вы загадали число: %v\n Так ли это?\n", number)
			fmt.Scan(&answer)
		}
	}

	if answer == success {
		fmt.Println("Ура! :)")
	} else {
		if answer == more {
			number += 1
			fmt.Printf("Вы загадали число: %v\n Так ли это?\n", number)
			fmt.Scan(&answer)
		} else {
			number -= 1
			fmt.Printf("Вы загадали число: %v\n Так ли это?\n", number)
			fmt.Scan(&answer)
		}
	}

	if answer == success {
		fmt.Println("Ура! :)")
	} else {
		if answer == more {
			number += 1
			fmt.Printf("Вы загадали число: %v\n Так ли это?\n", number)
			fmt.Scan(&answer)
		} else {
			number -= 1
			fmt.Printf("Вы загадали число: %v\n Так ли это?\n", number)
			fmt.Scan(&answer)
		}
		if answer == success {
			fmt.Println("Ура! :)")
		} else {
			fmt.Println("К сожалению мне не удалось угадать ваше число с 4-х попыток :(")
		}
	}
}
