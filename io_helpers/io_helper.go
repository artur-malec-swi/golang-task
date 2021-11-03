package io_helpers

import (
	"encoding/json"
	"os"
	"task/parsers"
	"task/text_analyzer"
)

func SaveByteArrayToFile(data []byte, filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	file.Write(data)
	defer file.Close()
	return nil
}

func SaveStringToFile(data string, filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	file.WriteString(data)
	defer file.Close()
	return nil
}

func MergeToFiles(filename1 string, filename2 string, resultFilename string) error {
	data1 := make(map[string]int)
	data2 := make(map[string]int)

	file1, err := os.ReadFile(filename1)

	if err != nil {
		return err
	}

	file2, err := os.ReadFile(filename2)
	if err != nil {
		return err
	}

	json.Unmarshal([]byte(file1), &data1)
	json.Unmarshal([]byte(file2), &data2)

	text_analyzer.MergeMaps(data1, data2)

	outputFile, err := os.Create(resultFilename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	byteArrOfJson, err := parsers.CreateJsonFromMap(data2)

	if err != nil {
		return err
	}

	SaveByteArrayToFile(byteArrOfJson, resultFilename)

	return nil
}
