package main

import "fmt"

func main() {

	var fistExam, secondExam, thirdExam int
	var pass = 275

	fmt.Println("Расчет проходного балла по результатам ЕГЭ")
	fmt.Printf("Введите оценку за первый экзамен: ")
	fmt.Scan(&fistExam)
	fmt.Printf("Введите оценку за первый экзамен: ")
	fmt.Scan(&secondExam)
	fmt.Printf("Введите оценку за первый экзамен: ")
	fmt.Scan(&thirdExam)

	if fistExam < 0 || secondExam < 0 || thirdExam < 0 || fistExam > 100 || secondExam > 100 || thirdExam > 100 {
		fmt.Println("Вы ввели некорректное значение оценки, количество баллов должно быть от 0 до 100")
	} else {
		if sum := fistExam + secondExam + thirdExam; sum < pass {
			fmt.Printf("По итогам за три экзамена вы набрали всего: %d, вам не хватило: %d \n", sum, pass-sum)
		} else {
			fmt.Printf("По итогам за три экзамена вы набрали: %d баллов, поздравляем! \n", sum)
		}
	}
}
