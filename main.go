package main

import (
	"errors"
	"fmt"
	"strings"
)

var nameError = errors.New("")

func getTopWords(wordMap map[string]int, n int) []string {
	a := []string{}
	fmt.Println("YOU")
	for word, _ := range wordMap {
		if n > 0 {
			a = append(a, word)
			n--
		} else {
			break
		}
	}
	return a
}

func AnalyzeText(text string) {

	text = strings.ReplaceAll(text, ".", " ")
	text = strings.ReplaceAll(text, "?", " ")
	text = strings.ReplaceAll(text, "!", " ")
	sts := strings.Split(text, " ")

	fmt.Println(fmt.Sprintf("Количество слов: %d", len(sts)))

	uniq := 0
	a := map[string]int{}
	for _, slv := range sts {
		if _, ok := a[slv]; ok {
			a[slv]++
		} else {
			uniq++
			a[slv] = 1
		}
	}
	fmt.Println(fmt.Sprintf("Количество уникальных слов: %d", uniq))
	tops := getTopWords(a, 5)
	fmt.Println(fmt.Sprintf("Самое часто встречающееся слово: %s (встречается %d раз)", tops[0], a[tops[0]]))
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for i := 0; i < 5; i++ {
		fmt.Println(fmt.Sprintf("%s: %d раз", tops[i], a[tops[i]]))
	}

}
