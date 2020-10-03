package main

import "fmt"

func main() {

	var students, groups, studentId int

	fmt.Println("Определение группы для указанного студента")
	fmt.Println("-----------------------------------")
	fmt.Printf("Укажите общее количество студентов: ")
	fmt.Scan(&students)
	fmt.Printf("Укажите общее количество групп: ")
	fmt.Scan(&groups)
	fmt.Printf("Укажите порядковый номер студента: ")
	fmt.Scan(&studentId)

	if students < 0 || groups < 0 {
		fmt.Println("Вы ввели некорректное количество студентов или групп")
	} else if students < groups {
		fmt.Println("Студентов не может быть меньше чем групп")
	} else if studentId > students {
		fmt.Println("Порядковый номер студента не может быть больше количества студентов")
	} else {
		studentsInGroup := students / groups
		currentGroup := studentId/studentsInGroup + 1
		fmt.Printf("Студент с порядковым номером %d прикреплен к группе %d\n", studentId, currentGroup)
	}
}
