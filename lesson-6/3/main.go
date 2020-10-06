package main

import "fmt"

func main() {

	var price, discont, totalDiscont int
	fmt.Printf("Введите стоимость товара (руб): ")
	fmt.Scan(&price)
	fmt.Printf("Введите скидку на товар: ")
	fmt.Scan(&discont)

	if discont > 30 {
		discont = 30
	}
	if totalDiscont = price * discont / 100; totalDiscont > 2000 {
		totalDiscont = 2000
	}

	fmt.Printf("Стоимость товара с учетом скидки составит: %d\n", price-totalDiscont)

}
