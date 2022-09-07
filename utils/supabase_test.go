package utils

import (
	"testing"
	"fmt"
)

func TestGetItemsForBusinessLocation(t *testing.T) {
	result, num, err := GetItemsForBusinessLocation("5f3f8f8f-f8f8-f8f8-f8f8-a0a0a0a0a0a0")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
	fmt.Println(num)
}