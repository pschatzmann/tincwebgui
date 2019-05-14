package service

/**
*  Formatters convert the result from a tinc call to the formatted result in
*  json (or text if there is no formatted information) which is returned by the webservice
**/
import (
	"encoding/json"
	"errors"
	"log"
	"strings"
)

// Mime - mime string: e.g. text/plain
type Mime string

// CommandFormatter - just converts the bytes stream to a string
func CommandFormatter(out []byte) ([]byte, Mime, error) {
	str := string(out[:])
	result := strings.TrimSpace(str)
	return []byte(result), "text/plain", nil
}

// JSONFormatter - just converts the bytes to a string and returns it as application/json
func JSONFormatter(out []byte) ([]byte, Mime, error) {
	str := string(out[:])
	result := strings.TrimSpace(str)
	return []byte(result), "application/json", nil
}

// ArrayFormatter - converts the result into an json array of strings
func ArrayFormatter(out []byte) ([]byte, Mime, error) {
	// split result string to array
	str := string(out[:])
	// handle new lines for windowns and linux
	str = strings.Replace(str, "\r\n", "\n", -1)
	// remove empty lines
	str = strings.TrimSpace(str)
	// convert to array
	var array = strings.Split(str, "\n")
	result, err := json.Marshal(array)
	if err != nil {
		return nil, "", err
	}

	// return result as json array
	return []byte(result), "application/json", nil
}

// Invitation - parsed invitation line
type Invitation struct {
	Invitation string
	Name       string
}

// InvitationsFormatter - converts the result into an json array of strings
func InvitationsFormatter(out []byte) ([]byte, Mime, error) {
	// split result string to array
	str := string(out[:])
	// handle new lines for windowns and linux
	str = strings.Replace(str, "\r\n", "\n", -1)
	// remove empty lines
	str = strings.TrimSpace(str)
	// convert to array
	var array = strings.Split(str, "\n")

	// build Invitations
	resultArray := make([]Invitation, 0)
	for _, line := range array {
		if line != "No outstanding invitations." {
			resultLine := strings.Fields(line)
			if len(resultLine) != 2 {
				return nil, "", errors.New(line)
			}
			invitation := Invitation{resultLine[0], resultLine[1]}
			resultArray = append(resultArray, invitation)
		}
	}

	// convert to Json
	result, err := json.Marshal(resultArray)
	if err != nil {
		return nil, "", err
	}

	// return result as json array
	return []byte(result), "application/json", nil
}

// NodesFormatter - converts the info result into an json object
// docker id 1d6c0791469a at unknown port unknown cipher 0 digest 0 maclength 0 compression 0 options 0 status 0800 nexthop - via - distance 0 pmtu 9018 (min 0 max 9018) rx 0 0 tx 0 0
func NodesFormatter(out []byte) ([]byte, Mime, error) {
	var nodeFields = []string{"name", "id", "at", "port", "cipher", "digest", "maclength", "compression", "options", "status", "nexthop", "distance", "pmtu", "rx", "tx"}
	return parsingFormatter(nodeFields, out)
}

// SubnetsFormatter - converts the info result into an json object
// ff:ff:ff:ff:ff:ff owner (broadcast)
// 255.255.255.255 owner (broadcast)
// 172.16.0.0/24 owner mcbookair
// 224.0.0.0/4 owner (broadcast)
// ff00::/8 owner (broadcast)
func SubnetsFormatter(out []byte) ([]byte, Mime, error) {
	var subnetFields = []string{"ip", "owner"}
	return parsingFormatter(subnetFields, out)
}

// ConnectionsFormatter - converts the info result into an json object
// dump connections
// <control> at localhost port unix options 0 socket 11 status 200
func ConnectionsFormatter(out []byte) ([]byte, Mime, error) {
	var connectionFields = []string{"control", "port", "options", "socket", "status"}
	return parsingFormatter(connectionFields, out)
}

// InfoFormatter - converts the info result into an json object
// info node (NODE|SUBNET|ADDRESS)
//tinc> info docker
//Node:         docker
//Node ID:      1d6c0791469a
//Address:      unknown port unknown
//Last seen:    never
//Status:
//Options:
//Protocol:     17.0
//Reachability: unreachable
//RX:           0 packets  0 bytes
//TX:           0 packets  0 bytes
//Edges:
//Subnets:
func InfoFormatter(out []byte) ([]byte, Mime, error) {
	// split result string to array
	str := string(out[:])
	// handle new lines for windowns and linux
	str = strings.Replace(str, "\r\n", "\n", -1)
	// convert to array
	var array = strings.Split(str, "\n")

	// build Invitations
	resultMap := map[string]string{}
	for _, line := range array {
		if line != "No outstanding invitations." {
			resultLine := strings.Split(line, ":")
			key := resultLine[0]
			// handle special value cases where empty or multiple :
			value := ""
			if len(resultLine) > 1 {
				value = strings.TrimSpace(strings.Join(resultLine[1:], " "))
			}
			// we just add informatioin which contain a value
			if value != "" {
				if key == "RX" || key == "TX" {
					packets, bytes := parseRXTX(value)
					resultMap[key+"-packets"] = packets
					resultMap[key+"-bytes"] = bytes
				} else {
					resultMap[key] = strings.TrimSpace(value)
				}
			}
		}
	}
	// convert to Json
	result, err := json.Marshal(resultMap)
	if err != nil {
		return nil, "", err
	}
	// return result as json array
	return []byte(result), "application/json", nil
}

// parseRXTX - parse a RX or TX string of the format  "0 packets  0 bytes"
func parseRXTX(in string) (string, string) {
	line := strings.ReplaceAll(in, "packets", ",")
	line = strings.ReplaceAll(line, "bytes", "")
	lines := strings.Split(line, ",")
	packets := strings.TrimSpace(lines[0])
	bytes := strings.TrimSpace(lines[1])
	return packets, bytes
}

// parsingFormatter - converts the result into an json array of strings based on the parsing fields
func parsingFormatter(fields []string, out []byte) ([]byte, Mime, error) {
	// split result string to array
	str := string(out[:])
	// handle new lines for windowns and linux
	str = strings.Replace(str, "\r\n", "\n", -1)
	// remove empty lines
	str = strings.TrimSpace(str)
	// convert to array
	var array = strings.Split(str, "\n")

	// build parsed result array
	resultArray := make([]map[string]string, 0)
	for _, line := range array {
		resultMap := parseLine(line, fields)
		resultArray = append(resultArray, resultMap)
	}

	// convert to Json
	result, err := json.Marshal(resultArray)
	if err != nil {
		return nil, "", err
	}

	// return result as json array
	return []byte(result), "application/json", nil
}

//parseLine - Simple parser which splits the line by attribute names
func parseLine(line string, attributes []string) map[string]string {
	// add start and end of line delimter
	lineEx := attributes[0] + " " + line + ";"
	attributes = append(attributes, ";")

	// process all attributes
	result := map[string]string{}
	for pos := range attributes[1:len(attributes)] {
		var startName = attributes[pos-1]
		var endName = attributes[pos]
		// split words
		lineArray := strings.Fields(lineEx)
		var startPos = find(lineArray, startName)
		var endPos = find(lineArray, endName)
		if startPos > endPos {
			// add value to map
			attributeValue := strings.Join(lineArray[startPos:endPos], " ")
			result[startName] = attributeValue
		} else {
			log.Println("Warning in parseLine - The field was ignored:", startName)
		}
	}
	return result
}

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}
