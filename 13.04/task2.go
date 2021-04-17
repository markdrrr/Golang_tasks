package main

import (
	"fmt"
	"strings"
)

/*
	Напишите функцию для подсчета частоты упоминания слов в строке текста
	и возвращения карты со словами и числом, указывающем, сколько раз они
	употребляются. Функция должна конвертировать текст в нижний регистр и
	обрезать знаки препинания. Используйте пакет strings. Функции, которые
	пригодятся для выполнения данного задания: Fields, ToLower и Trim.
*/
func main() {
	text := "Lorem Ipsum - это текст рыба, часто используемый в печати и вэб-дизайне. Lorem Ipsum является стандартной рыбой для текстов на латинице с начала XVI века. В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, используя Lorem Ipsum для распечатки образцов. Lorem Ipsum не только успешно пережил без заметных изменений пять веков, но и перешагнул в электронный дизайн. Его популяризации в новое время послужили публикация листов Letraset с образцами Lorem Ipsum в 60-х годах и, в более недавнее время, программы электронной вёрстки типа Aldus PageMaker, в шаблонах которых используется Lorem Ipsum."
	countWords(text)
}

func countWords(text string) map[string]int {
	result := make(map[string]int)
	text2 := strings.ToLower(text)
	text_slice := strings.Fields(text2)
	for _, v := range text_slice {
		value := strings.Trim(v, "-,.")
		_, ok := result[value]
		if !ok {
			result[value] = 1
		}
		if ok {
			result[value] += 1
		}
	}
	for k, v := range result {
		fmt.Println(k, ":", v)
	}
	return result
}
