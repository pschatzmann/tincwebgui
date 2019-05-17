package service

/**
 * Collects the netork px and rx statistics so that we can display it in a chart
 **/

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// record for last measurement
type measurementsRecordType struct {
	timestamp time.Time
	tx        measurementXStruct
	rx        measurementXStruct
}

// values for last measurement
type measurementXStruct struct {
	bytes   int64
	packets int64
}

// MeasurementsPerSecondType - Indivudual Record for the Recoding Result per node
type MeasurementsPerSecondType struct {
	Timestamp time.Time
	Seconds   float64
	Bytes     resultPerSecond
	Packets   resultPerSecond
}

type resultPerSecond struct {
	Rx float64
	Tx float64
}

// record of last measuremnets for each node
var lastMeasurements = make(map[string]*measurementsRecordType)

// chanel for timer
var quit = make(chan int)
var STOP = 0

// timer for the recoring of data
var ticker *time.Ticker

// measurement database
var measurementResult = make(map[string][]MeasurementsPerSecondType)

// collect data at point of time
func measure() {
	log.Println("measure")
	out, err := exec.Command("tinc", "dump", "nodes").CombinedOutput()
	if err != nil {
		log.Println("measure with error", err)
		return
	}
	timestamp := time.Now()
	dataList := Parse(nodeFields, string(out))
	for _, dataMap := range dataList {
		// process values for each node
		name := dataMap["name"]
		rxPackets, rxBytes := parseX(dataMap["rx"])
		txPackets, txBytes := parseX(dataMap["tx"])
		newRec := measurementsRecordType{}
		newRec.timestamp = timestamp
		newRec.rx.bytes = rxBytes
		newRec.rx.packets = rxPackets
		newRec.tx.bytes = txBytes
		newRec.tx.packets = txPackets

		// extend result database
		addResultValues(name, &newRec)
	}
	return
}

func parseX(line string) (int64, int64) {
	sa := strings.Split(line, " ")
	if len(sa) == 2 {
		packets, _ := strconv.ParseInt(sa[0], 10, 64)
		bytes, _ := strconv.ParseInt(sa[1], 10, 64)
		return packets, bytes
	}
	return 0, 0
}

func addResultValues(name string, measurement *measurementsRecordType) {
	priorMeasurement := lastMeasurements[name]
	if priorMeasurement != nil {
		duration := measurement.timestamp.Sub(priorMeasurement.timestamp)
		// calculate rates
		rec := MeasurementsPerSecondType{}
		rec.Timestamp = measurement.timestamp
		rec.Seconds = duration.Seconds()
		rec.Bytes.Rx = float64(measurement.rx.bytes-priorMeasurement.rx.bytes) / duration.Seconds()
		rec.Packets.Rx = float64(measurement.rx.packets-priorMeasurement.rx.packets) / duration.Seconds()
		rec.Bytes.Tx = float64(measurement.tx.bytes-priorMeasurement.tx.bytes) / duration.Seconds()
		rec.Packets.Tx = float64(measurement.tx.packets-priorMeasurement.tx.packets) / duration.Seconds()
		// append record
		measurementResult[name] = append(measurementResult[name], rec)
	}
	// update last record so that we calculate the rates
	lastMeasurements[name] = measurement
}

// StartNetworkTraffic - starts to record the network tx and rx
func StartNetworkTraffic() {
	if ticker == nil {
		ticker = time.NewTicker(5 * time.Second)
		// we restart with clean data
		lastMeasurements = make(map[string]*measurementsRecordType)
		measurementResult = make(map[string][]MeasurementsPerSecondType)
		go func() {
			for {
				select {
				case <-ticker.C:
					measure()
				case <-quit:
					log.Println("Stopping Recording...")
					if ticker != nil {
						ticker.Stop()
					}
					ticker = nil
					return
				}
			}
		}()
	} else {
		log.Println("The traffic recording is already active")
	}
}

// StopNetworkTraffic - stops the network tx and rx recording
func StopNetworkTraffic() {
	quit <- STOP
}

// GetNetworkTraffic - Privdes the recorded network statistics
func GetNetworkTraffic() map[string][]MeasurementsPerSecondType {
	return measurementResult
}

// GetNetworkTrafficInJSON - Provides the recorded network statistics as JSON
func GetNetworkTrafficInJSON() []byte {
	result := GetNetworkTraffic()
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return nil
	}
	// return result as json array
	return resultJSON
}

// NetworkTrafficHandler - Provides the recorded network statistics as JSON to http
func NetworkTrafficHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("NetworkTrafficHandler")
	w.Header().Set("Content-Type", "application/json")
	w.Write(GetNetworkTrafficInJSON())
}

// NetworkTrafficStartHandler - start recording
func NetworkTrafficStartHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("NetworkTrafficStartHandler")
	StartNetworkTraffic()
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}

// NetworkTrafficStopHandler - stops recording
func NetworkTrafficStopHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("NetworkTrafficStopHandler")
	StopNetworkTraffic()
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}
