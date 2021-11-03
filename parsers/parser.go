package parsers

import (
	"encoding/json"
)

func CreateJsonFromMap(data map[string]int) ([]byte, error) {
	byteArr, err := json.Marshal(data)

	if err != nil {
		return []byte{}, err
	}

	return byteArr, nil
}
