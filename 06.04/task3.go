package main

import "fmt"

/*
	Задача. Переменная num может принимать 4 значения:
	1, 2, 3 или 4. Если она имеет значение '1', то в переменную result
	запишем 'зима', если имеет значение '2' – 'весна' и так далее.
	Задачу выполнить с использованием функции, которая просто будет
	возвращать результат (сезон)
*/

func main() {
	var num int
	fmt.Println("Ввеидте цифру от 1 до 4:")
	if num < 1 || num > 4 {
		fmt.Println("Неверный ввод")
		return
	}
	fmt.Scan(&num)
	fmt.Println(season(num))
}

func season(num int) string {
	var result string
	if num == 1 {
		result = "Зима"
	} else if num == 2 {
		result = "Весна"
	} else if num == 3 {
		result = "Лето"
	} else if num == 4 {
		result = "Осень"
	}
	return result
}
