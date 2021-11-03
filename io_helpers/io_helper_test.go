package io_helpers

import (
	"os"
	"testing"
)

func TestMergeTwoFiles(t *testing.T) {
	filename1 := "file1.json"
	filename2 := "file2.json"
	resultFilename := "file3.json"
	firstFileContent := "{\"a\":8,\"ac\":8,\"adipiscing\":7}"
	secondFileContent := "{\"a\":8,\"ac\":8,\"abc\":7}"
	expectedResult := "{\"a\":16,\"abc\":7,\"ac\":16,\"adipiscing\":7}"

	file1, err := os.Create(filename1)

	if err != nil {
		t.Error("Failed on creating the test file.")
	}
	file1.WriteString(firstFileContent)
	file1.Close()

	file2, err := os.Create(filename2)

	if err != nil {
		t.Error("Failed on creating the test file.")
	}
	file2.WriteString(secondFileContent)
	file2.Close()

	err = MergeToFiles(filename1, filename2, resultFilename)

	if err != nil {
		t.Error("Failed on merging the file.")
	}

	resultFileContent, err := os.ReadFile(resultFilename)

	if err != nil {
		t.Error("Failed on reading the result file.")
	}
	if expectedResult != string(resultFileContent) {
		t.Error("Test failed. Expected result: {}, received: {}", expectedResult, string(resultFileContent))
	}

	defer os.Remove(filename1)
	defer os.Remove(filename2)
	defer os.Remove(resultFilename)
}
