package main

import "testing"

func TestSub(t *testing.T) {
	result := sub(5, 2)

	if result != 3 {
		t.Error(ERROR_MESSAGE + "3")
	}
}

func TestSubFailed(t *testing.T) {
	result := sub(5, 5)

	if result == 0 {
		t.Error(TEST_FAILED)
	}
}
