package main

import "fmt"

/*
	Задача. Вывести на экран ряд чисел Фибоначчи, состоящий из n элементов.
	Числа Фибоначчи – это элементы числовой последовательности
	0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, …,
	в которой каждое последующее число равно сумме двух предыдущих.
*/

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) []int {
	var numbers = []int{0, 1}
	var sum int
	fib_1, fib_2 := 0, 1
	for i := 0; i < n; i++ {
		sum = fib_1 + fib_2
		fib_1 = fib_2
		fib_2 = sum
		numbers = append(numbers, fib_2)
	}
	return numbers
}
