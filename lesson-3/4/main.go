package main

import "fmt"

func main() {

	{
		fmt.Println("1. Расчет высоты бамбука на третий день")

		var rootBamboo = 100
		var growBamboo = 50
		var eatBamboo = 20

		fmt.Printf("Бамбук вырастет на %dсм к середине трерьго дня\n", rootBamboo+growBamboo-eatBamboo+growBamboo-eatBamboo+growBamboo/2)
		fmt.Printf("Бамбук вырастет на %dсм к середине трерьго дня\n", rootBamboo+(growBamboo-eatBamboo)*2+growBamboo/2)
		fmt.Println("-------------")
		fmt.Println("2. Расчет количества дней на продажу бамбука")
		fmt.Printf("Требуется %d дней чтобы бамбук вырос на 3 метра\n", (300-rootBamboo)/(growBamboo-eatBamboo))
		fmt.Println("-------------")

	}

	fmt.Println("3. Расчет роста пользовательского бамбука")

	var rootBamboo, growBamboo, eatBamboo, days, heightBamboo int

	fmt.Printf("Введите высоту саженца бамбука (см): ")
	fmt.Scan(&rootBamboo)
	fmt.Printf("Введите скорость роста бамбука за день (см): ")
	fmt.Scan(&growBamboo)
	fmt.Printf("Введите скорость поедания бамбука гусеницами (см): ")
	fmt.Scan(&eatBamboo)
	fmt.Printf("Задайте количество дней выращивния: ")
	fmt.Scan(&days)
	fmt.Printf("Задайте необходимую высоту бамбука (см): ")
	fmt.Scan(&heightBamboo)
	fmt.Printf("\n\nБамбук вырастет на %dсм за %d дней\n", rootBamboo+growBamboo*days-eatBamboo*(days-1), days)
	fmt.Printf("\nТребуется %d дней чтобы бамбук вырос на %d сантиметров\n", (heightBamboo-rootBamboo)/(growBamboo-eatBamboo), heightBamboo)
}
