package main

import (
	"fmt"
	"slices"
	"strings"
)

func getTopWords(wordMap map[string]int, n int) []string{
	result := make([]string, n)
	for i := 0; i < n; i++{
		var res_word string
		for word, count := range wordMap {
			if !slices.Contains(result, word) && count > wordMap[res_word]{
				res_word = word
			}
		}
		result[i] = res_word
	}

	return result
}

func AnalyzeText(text string){
	text = strings.Replace(text, ".", "", -1)
	text = strings.Replace(text, "!", "", -1)
	text = strings.Replace(text, "?", "", -1)
	text = strings.Replace(text, ",", "", -1)

	words := strings.Split(text, " ")
	count_words := len(words)

	m := make(map[string]int)
	for _, word := range words{
		m[strings.ToLower(word)] += 1
	}
	count_unique_words := len(m)
	top := getTopWords(m, 5)
	
	fmt.Printf("Количество слов: %d\n", count_words)
	fmt.Printf("Количество уникальных слов: %d\n", count_unique_words)
	fmt.Printf(`Самое часто встречающееся слово: "%s" (встречается %d раз)
`,top[0], m[top[0]])
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for i := 0; i < 5; i++ {
		fmt.Printf(`"%s": %d раз
`, top[i], m[top[i]])
	}

}