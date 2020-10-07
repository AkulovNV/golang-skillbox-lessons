package main

import "fmt"

func main() {

	var fistBacketSize, secondBacketSize, thirdBacketSize int

	fmt.Println("Наполнение корзин разной емкости яблоками")
	fmt.Println("Укажите сколько яблок может уместиться в каждую из корзин")

	for {

		fmt.Printf("Емкость первой корзины: ")
		fmt.Scan(&fistBacketSize)
		fmt.Printf("Емкость второй корзины: ")
		fmt.Scan(&secondBacketSize)
		fmt.Printf("Емкость третьей корзины: ")
		fmt.Scan(&thirdBacketSize)

		if fistBacketSize < 0 || secondBacketSize < 0 || thirdBacketSize < 0 {
			fmt.Println("Емкость не может быть отрицательной, повторите ввод")
			continue
		} else {
			break
		}

	}
	putFistBacket, putSeconfBacket, putThirdBacket := 0, 0, 0
	maxFistSize, maxSecondSize, maxThirdSize := 0, 0, 0

	for {

		if putFistBacket == fistBacketSize && putSeconfBacket == secondBacketSize && putThirdBacket == thirdBacketSize {

			// Определяем последовательность заполнения корзин
			// Так и не пришло в голову как более лаконичнее и красивее эту проверку сделать...
			if maxFistSize > maxSecondSize && maxFistSize > maxThirdSize {
				if maxSecondSize > maxThirdSize {
					fmt.Println("Емкость первой корзины заполнена")
					fmt.Println("Емкость второй корзины заполнена")
					fmt.Println("Емкость третье корзины заполнена")
				} else {
					fmt.Println("Емкость первой корзины заполнена")
					fmt.Println("Емкость третье корзины заполнена")
					fmt.Println("Емкость второй корзины заполнена")
				}
			} else if maxSecondSize > maxFistSize && maxSecondSize > maxThirdSize {
				if maxFistSize > maxThirdSize {
					fmt.Println("Емкость второй корзины заполнена")
					fmt.Println("Емкость первой корзины заполнена")
					fmt.Println("Емкость третье корзины заполнена")
				} else {
					fmt.Println("Емкость второй корзины заполнена")
					fmt.Println("Емкость третье корзины заполнена")
					fmt.Println("Емкость первой корзины заполнена")
				}
			} else if maxThirdSize > maxFistSize && maxThirdSize > maxSecondSize {
				if maxFistSize > maxSecondSize {
					fmt.Println("Емкость третье корзины заполнена")
					fmt.Println("Емкость первой корзины заполнена")
					fmt.Println("Емкость второй корзины заполнена")
				} else {
					fmt.Println("Емкость третье корзины заполнена")
					fmt.Println("Емкость второй корзины заполнена")
					fmt.Println("Емкость первой корзины заполнена")
				}
			} else if maxFistSize == maxSecondSize {
				fmt.Println("Емкость первой корзины заполнена")
				fmt.Println("Емкость второй корзины заполнена")
				fmt.Println("Емкость третье корзины заполнена")
			} else if maxFistSize == maxThirdSize {
				fmt.Println("Емкость первой корзины заполнена")
				fmt.Println("Емкость третье корзины заполнена")
				fmt.Println("Емкость второй корзины заполнена")
			} else if maxFistSize == maxThirdSize {
				fmt.Println("Емкость второй корзины заполнена")
				fmt.Println("Емкость третье корзины заполнена")
				fmt.Println("Емкость первой корзины заполнена")
			} else {
				fmt.Println("Емкость первой корзины заполнена")
				fmt.Println("Емкость второй корзины заполнена")
				fmt.Println("Емкость третье корзины заполнена")
			}

			fmt.Println("Все корзины заполнены")
			break
		}

		if putFistBacket < fistBacketSize {
			putFistBacket += 1
		} else {
			maxFistSize += 1
		}

		if putSeconfBacket < secondBacketSize {
			putSeconfBacket += 1
		} else {
			maxSecondSize += 1
		}

		if putThirdBacket < thirdBacketSize {
			putThirdBacket += 1
		} else {
			maxThirdSize += 1
		}
	}

}
