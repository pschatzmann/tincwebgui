package service

/**
*  Parser for tinc dump commands: It Parses a string with the help of an array
*  of delimiter fields into a Map
**/
import (
	"log"
	"strings"
)

// Parse - from string to map[string]string
func Parse(fields []string, input string) []map[string]string {
	// handle new lines for windowns and linux
	str := strings.Replace(input, "\r\n", "\n", -1)
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
	return resultArray
}

//parseLine - Simple parser which splits the line by attribute names
func parseLine(line string, attributes []string) map[string]string {
	// add start and end of line delimter
	lineEx := attributes[0] + " " + line + " ;"
	attributes = append(attributes, ";")

	// process all attributes
	result := map[string]string{}
	for pos := 0; pos < len(attributes)-1; pos++ {
		var startName = attributes[pos]
		var endName = attributes[pos+1]
		// split words
		lineArray := strings.Fields(lineEx)
		var startPos = find(lineArray, startName, pos) + 1
		var endPos = find(lineArray, endName, pos)
		if startPos < endPos {
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
func find(a []string, x string, start int) int {
	for i, n := range a {
		if i >= start {
			if x == n {
				return i
			}
		}
	}
	return -1
}
