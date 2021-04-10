package main

import "fmt"

/*
	Задача. В переменной min лежит число от 0 до 59.
	Определите в какую четверть часа попадает это число
	(в первую, вторую, третью или четвертую).
*/

func main() {
	var number int
	fmt.Print("Введите число от 0 - 59: ")
	fmt.Scan(&number)

	if number < 1 || number > 59 {
		fmt.Println("Неверный ввод")
		return
	}
	if number <= 15 {
		fmt.Println("первая четверть")
	} else if number <= 30 {
		fmt.Println("Вторая четверть")
	} else if number <= 45 {
		fmt.Println("Третья четверть")
	} else if number <= 59 {
		fmt.Println("Четвертая четверть")
	}
}
