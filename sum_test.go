package main

import "testing"

const ERROR_MESSAGE = "The result must be "

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error(ERROR_MESSAGE + "5")
	}
}


