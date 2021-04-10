package main

import "fmt"

/*
	Задача. Если переменная a равна или меньше 1, а переменная b больше
	или равна 3, то выведите сумму этих переменных, иначе выведите их разность
	(результат вычитания). Проверьте работу скрипта при a и b, равном 1 и 3, 0 и 6, 3 и 5.
*/

func main() {
	var a, b int
	fmt.Println("Input a")
	fmt.Scan(&a)
	fmt.Println("Input b")
	fmt.Scan(&b)

	if (a <= 1) && (b >= 3) {
		a = a + b
		fmt.Println(a)
	} else {
		a = a - b
		fmt.Println(a)
	}
}
