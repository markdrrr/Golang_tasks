package main

import "fmt"

/*
	Есть слайс из 1000 элементов типа int,
	нужно его забить и отсортировать в обратном порядке
	без использования встроенных функций
*/
func main() {
	var elements []int
	for i := 0; i < 1000; i++ {
		elements = append(elements, i)
	}
	for i, j := 0, len(elements)-1; i < j; i, j = i+1, j-1 {
		elements[i], elements[j] = elements[j], elements[i]
	}
	fmt.Println(elements)
}
