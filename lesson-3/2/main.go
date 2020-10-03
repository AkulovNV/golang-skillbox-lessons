package main

import "fmt"

func main() {

	fmt.Println("Симулятор маршрутки")

	var passengersIn, passengersEnter, passengersLeft, passengersTotal int
	var ticketCost = 20

	fmt.Println("Прибываем на остановку Улица Программистов. В салоне пассажиров: 0")
	fmt.Printf("Сколько пассажиров вышло на остановке? ")
	fmt.Scan(&passengersLeft)
	fmt.Printf("Сколько пассажиров зашло на остановке? ")
	fmt.Scan(&passengersEnter)
	passengersIn += passengersEnter - passengersLeft
	passengersTotal += passengersEnter
	fmt.Printf("Отправляемся с остановки Улица Программистов. В салоне пассажиров: %d\n", passengersIn)
	fmt.Println("-----------Едем---------")

	fmt.Printf("Прибываем на остановку Проспект Алгоритмов. В салоне пассажиров: %d\n", passengersIn)
	fmt.Printf("Сколько пассажиров вышло на остановке? ")
	fmt.Scan(&passengersLeft)
	fmt.Printf("Сколько пассажиров зашло на остановке? ")
	fmt.Scan(&passengersEnter)
	passengersIn += passengersEnter - passengersLeft
	passengersTotal += passengersEnter
	fmt.Printf("Отправляемся с остановки Проспект Алгоритмов. В салоне пассажиров: %d\n", passengersIn)
	fmt.Println("-----------Едем---------")

	fmt.Printf("Прибываем на остановку Площадь Рекурсии. В салоне пассажиров: %d\n", passengersIn)
	fmt.Printf("Сколько пассажиров вышло на остановке? ")
	fmt.Scan(&passengersLeft)
	fmt.Printf("Сколько пассажиров зашло на остановке? ")
	fmt.Scan(&passengersEnter)
	passengersIn += passengersEnter - passengersLeft
	passengersTotal += passengersEnter
	fmt.Printf("Отправляемся с остановки Площадь Рекурсии. В салоне пассажиров: %d\n", passengersIn)
	fmt.Println("-----------Едем---------")

	fmt.Printf("Прибываем на остановку Бульвар Исключений. В салоне пассажиров: %d\n", passengersIn)
	fmt.Printf("Сколько пассажиров вышло на остановке? ")
	fmt.Scan(&passengersLeft)
	fmt.Printf("Сколько пассажиров зашло на остановке? ")
	fmt.Scan(&passengersEnter)
	passengersIn += passengersEnter - passengersLeft
	passengersTotal += passengersEnter
	fmt.Printf("Отправляемся с остановки Бульвар Исключений. В салоне пассажиров: %d\n", passengersIn)
	fmt.Println("-----------Едем---------")

	fmt.Println("Прибываем на конечную остановку")
	fmt.Println("-------------------------------")

	var receipt = ticketCost * passengersTotal
	var salary = receipt / 4
	var gas = receipt / 5
	var tax = receipt / 5
	var repair = receipt / 5
	var totalIncome = receipt - salary - gas - tax - repair

	fmt.Printf("Всего заработали: %d рублей\n", receipt)
	fmt.Printf("Зарплата водителя: %d рублей\n", salary)
	fmt.Printf("Расходы на топливо: %d рублей\n", gas)
	fmt.Printf("Налоги: %d рублей\n", tax)
	fmt.Printf("Расходы на ремонт машины: %d рублей\n", repair)
	fmt.Printf("Итого доход: %d рублей\n", totalIncome)
}
