package utils

import (
	"testing"
)

func TestPrint1(t *testing.T) {
	IPCount = 0
	ip := "localhost:8888"
	PrintIP(&ip)
	if IPCount == 0 {
		t.Errorf("The count was %v", IPCount)

	}
}

func TestPrintN(t *testing.T) {
	IPCount = 0
	ip := "0.0.0.0:8888"
	PrintIP(&ip)
	if IPCount == 0 {
		t.Errorf("The count was %v", IPCount)
	}
}
