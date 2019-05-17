package service

import (
	"log"
	"testing"
	"time"
)

// TestMeasure - test recording of network statistics
func TestMeasure(t *testing.T) {
	StartNetworkTraffic()
	time.Sleep(12 * time.Second)
	StopNetworkTraffic()
	speed := GetNetworkTraffic()

	if len(speed) == 0 {
		t.Errorf("The network statistics are empty")
	}

	StopNetworkTraffic()

	log.Println(GetNetworkTrafficInJSON())
}
