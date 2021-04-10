package main

import "fmt"

/*
	Задача. Вычислить факториал введенного числа.
	Факториалом числа называют произведение всех натуральных чисел
	до этого числа включительно. Например, факториал числа 4 равен 1*2*3*4 = 24.
*/

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(factorial(n))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return factorial(n-1) * n
}
