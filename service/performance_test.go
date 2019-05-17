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

	speed := GetNetworkTraffic()
	log.Println(GetNetworkTrafficInJSON(true, true))

	if len(speed) == 0 {
		t.Errorf("The network statistics are empty")
	}

	for key, value := range speed {
		if len(value) == 0 {
			t.Errorf("The network statistics are empty for " + key)
		}
	}

	StopNetworkTraffic()
}
