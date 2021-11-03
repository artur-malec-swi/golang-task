package text_analyzer

import (
	"reflect"
	"testing"
)

func TestGetMostCommonWords(t *testing.T) {
	longString := "test test aa bb zzz kkk ll oo aa test"
	expectedResult := map[string]int {"test": 3, "aa": 2}

	result, err := GetMostCommonWords(longString, 2)
	if err != nil {
		t.Error("Test failed. Error message: {}", err)
	}

	eq := reflect.DeepEqual(result, expectedResult)
	if !eq {
		t.Error("Test failed. {} inputted, {} expected, {} received", longString, expectedResult, result)
	}
}
