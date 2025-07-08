package main

import "testing"

func TestSumX(t *testing.T) {
	result := sumX(3, 2)

	if result != 8 {
		t.Error(ERROR_MESSAGE + "8")
	}
}

func TestSumXfailed(t *testing.T) {
	result := sumX(3, 2)

	if result != 7 {
		t.Error(TEST_FAILED)
	}
}
