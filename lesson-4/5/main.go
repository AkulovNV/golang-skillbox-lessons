package main

import "fmt"

func main() {

	var dayOfWeek, guests, checkAmount int
	var mondayDiscont = 10
	var fridayDiscont = 5
	var service = 10

	fmt.Println("Расчет итоговой стомости заказа в ресторане")
	fmt.Printf("День недели (число): ")
	fmt.Scan(&dayOfWeek)
	fmt.Printf("Количество гостей: ")
	fmt.Scan(&guests)
	fmt.Printf("Общая сумма заказа: ")
	fmt.Scan(&checkAmount)

	if guests > 5 {
		checkAmount += checkAmount * service / 100
	}

	if dayOfWeek == 1 {
		fmt.Printf("Итого к оплате: %d\n", checkAmount-checkAmount*mondayDiscont/100)
	} else if dayOfWeek == 5 {
		fmt.Printf("Итого к оплате: %d\n", checkAmount-checkAmount*fridayDiscont/100)
	} else {
		fmt.Printf("Итого к оплате: %d\n", checkAmount)
	}

}
