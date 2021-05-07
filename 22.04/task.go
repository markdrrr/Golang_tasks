package main

import (
	"fmt"
	"strings"
	"sync"
)

/*
	Напишите функцию для подсчета частоты упоминания слов в строке текста
	и возвращения карты со словами и числом, указывающем, сколько раз они
	употребляются. Функция должна конвертировать текст в нижний регистр и
	обрезать знаки препинания. Используйте пакет strings. Функции, которые
	пригодятся для выполнения данного задания: Fields, ToLower и Trim.
*/
func main() {
	text := `Lorem Ipsum - это текст рыба, часто используемый в печати и вэб-дизайне. 
			Lorem Ipsum является стандартной рыбой для текстов на латинице с начала XVI века. 
			В то время некий безымянный печатник создал большую коллекцию размеров и форм шрифтов, 
			используя Lorem Ipsum для распечатки образцов. Lorem Ipsum не только успешно пережил 
			без заметных изменений пять веков, но и перешагнул в электронный дизайн. Его популяризации 
			в новое время послужили публикация листов Letraset с образцами Lorem Ipsum в 60-х годах и, 
			в более недавнее время, программы электронной вёрстки типа Aldus PageMaker, 
			в шаблонах которых используется Lorem Ipsum.`
	res := countWords(strings.ToLower(text))
	fmt.Println(res)
}

func countWords(text string) map[string]int {
	result := make(map[string]int)
	textSlice := strings.Fields(text)
	N := 3
	wg := sync.WaitGroup{}
	mux := sync.Mutex{}
	wg.Add(N)
	a := len(textSlice)
	start := 0
	end := int(len(textSlice)/N) + 1
	for i := 0; i < 3; i++ {
		go func(start int, end int) {
			getCountWord(textSlice[start:end], result, &mux, &wg)
		}(start, end)
		start = end
		end = end + int(len(textSlice)/3)
		if end > a {
			end = len(textSlice)
		}
	}
	wg.Wait()
	return result
}

func getCountWord(text []string, result map[string]int, mux *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range text {
		value := strings.Trim(v, "-,.")
		mux.Lock()
		result[value]++
		mux.Unlock()
	}
}
