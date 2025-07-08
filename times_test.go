package main

import "testing"

const TEST_FAILED = "Test failed"

func TestTimes(t *testing.T) {
	result := times(6, 2)

	if result != 12 {
		t.Error(ERROR_MESSAGE + "12")
	}
}

func TestTimesFail(t *testing.T) {
	result := times(6, 5)

	if result == 25 {
		t.Error(TEST_FAILED)
	}
}
