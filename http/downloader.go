package http

import (
	"io/ioutil"
	"net/http"
)

func GetHttpBody(url string, c chan string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		c <- string(bodyBytes)
	}
	return nil
}
