package text_analyzer

import (
	"errors"
	"sort"
	"strings"
)

func GetMostCommonWords(text string, numberOfItems int) (map[string]int, error) {
	wordsCountMap := make(map[string]int)

	if text == "" {
		return wordsCountMap, errors.New("Given string is empty")
	}

	words := strings.Split(text, " ")

	for _, word := range words {
		_, ok := wordsCountMap[word]
		if !ok {
			wordsCountMap[word] = 1
		} else {
			wordsCountMap[word]++
		}
	}

	keys := make([]string, 0, len(wordsCountMap))
	for key := range wordsCountMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return wordsCountMap[keys[i]] > wordsCountMap[keys[j]]
	})

	result := make(map[string]int)

	for _, key := range keys {
		result[key] = wordsCountMap[key]
		numberOfItems--
		if numberOfItems == 0 {
			break
		}
	}

	return result, nil
}

func MergeMaps(source map[string]int, target map[string]int) {
	for k, v := range source {
		_, ok := target[k]
		if !ok {
			target[k] = v
		} else {
			target[k] += v
		}
	}
}
