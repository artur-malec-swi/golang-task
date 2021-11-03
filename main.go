package main

import (
	"fmt"
	"task/http"
	"task/io_helpers"
	"task/parsers"
	"task/text_analyzer"
)

func main() {
	const firstUrl = "https://pastebin.com/raw/v0Sm2sfn"
	const secondUrl = "https://pastebin.com/raw/fysHJ7YX"
	const numberOfMostCommonItems = 50

	c1 := make(chan string)
	c2 := make(chan string)

	go http.GetHttpBody(firstUrl, c1)
	go http.GetHttpBody(secondUrl, c2)

	firstData, secondData := <-c1, <-c2

	io_helpers.SaveStringToFile(firstData, "out1.txt")
	io_helpers.SaveStringToFile(secondData, "out2.txt")

	mostCommonWordsForFirstData, e := text_analyzer.GetMostCommonWords(firstData, numberOfMostCommonItems)

	if e != nil {
		panic(e)
	}

	mostCommonWordsForSecondData, e := text_analyzer.GetMostCommonWords(secondData, numberOfMostCommonItems)

	data, _ := parsers.CreateJsonFromMap(mostCommonWordsForFirstData)

	io_helpers.SaveByteArrayToFile(data, "out1.json")

	data, _ = parsers.CreateJsonFromMap(mostCommonWordsForSecondData)

	io_helpers.SaveByteArrayToFile(data, "out2.json")

	io_helpers.MergeToFiles("out1.json", "out2.json", "out3.json")

	if e != nil {
		panic(e)
	}

	fmt.Println("Done")
}
