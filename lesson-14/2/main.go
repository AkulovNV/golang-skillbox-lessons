/*
Написать программу, которая с помощью функции генерирует 3 случайные точки в двумерном пространстве (две координаты),
а затем с помощью другой функции преобразует эти координаты по формулам: x1 = 2 * x + 10, y1 = -3 * y - 5.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandom(n int) int {
	return rand.Intn(n)
}

func getRandomPoint(p1, p2 int) (int, int) {
	return rand.Intn(p1), rand.Intn(p2)
}

func coordination(x, y int) string {
	x1 := 2*x + 10
	y1 := -3*y - 5
	return fmt.Sprintf("%d,%d", x1, y1)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	x1, y1 := getRandomPoint(10, 100)
	x2, y2 := getRandomPoint(10, 100)
	x3, y3 := getRandomPoint(10, 100)
	fmt.Println("Коордираны точки x1,y1: ", coordination(x1, y1))
	fmt.Println("Коордираны точки x2,y2: ", coordination(x2, y2))
	fmt.Println("Коордираны точки x3,y3: ", coordination(x3, y3))
}
