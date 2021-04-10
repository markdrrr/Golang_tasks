package main

import "fmt"

/*
	Задача. Напишите программу, доказывающую или проверяющую,
	что для множества натуральных чисел выполняется равенство:
	1+2+...+n = n(n+1)/2, где n - любое натуральное число.
*/

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(nature(n))
}

func nature(n int) bool {
	var resLeftPart, resRightPart int
	for i := 0; i <= n; i++ {
		resLeftPart += i
	}
	resRightPart = n * (n + 1) / 2
	fmt.Println(resLeftPart)
	fmt.Println(resRightPart)
	return resLeftPart == resRightPart
}
