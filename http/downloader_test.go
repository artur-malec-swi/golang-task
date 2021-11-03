package http

import (
	"testing"
)

func TestDownloader(t *testing.T) {
	url := "https://pastebin.com/raw/v0Sm2sfn"

	c := make(chan string)

	go GetHttpBody(url, c)

	result := <-c

	if result == "" {
		t.Error("Test failed. Got empty string.")
	}
}
