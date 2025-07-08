package main

import "testing"

const ERROR_MESSAGE = "The result must be "

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error(ERROR_MESSAGE + "5")
	}
}

func TestSub(t *testing.T) {
	result := sub(5, 2)

	if result != 3 {
		t.Error(ERROR_MESSAGE + "3")
	}
}

func TestTimes(t *testing.T) {
	result := times(6, 2)

	if result != 12 {
		t.Error(ERROR_MESSAGE + "12")
	}
}

func TestSumX(t *testing.T) {
	result := sumX(3, 2)

	if result != 8 {
		t.Error(ERROR_MESSAGE + "8")
	}
}
