package main

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	startToken     = "BENCHMARK"
	startTokenLen  = 9
	endToken       = "END_BENCH"
	separator      = ".."
	leftIndicator  = "["
	rightIndicator = "]"
	timeLayout     = "2006-01-02 15:04:05.000000 MST"
)

// the function below pares the log string by format below.
// "[time] [module] <optional content> BENCHMARK key_1:value_1..key_2:value_2.. ..key_n:value_n END_BENCH"
func Parse(input string) (*LogData, error) {
	var logData LogData
	logData.module = "undefined"
	var lIndex = strings.Index(input, leftIndicator)
	var rIndex = strings.Index(input, rightIndicator)
	if lIndex >= 0 && rIndex > 0 && lIndex < rIndex {
		var timeString = strings.TrimSpace(input[lIndex+1 : rIndex])
		var t, err = time.Parse(timeLayout, timeString)
		if err != nil {
			logData.timStamp = t
		}

		if rIndex+1 < len(input) {
			var content = input[rIndex+1:]
			lIndex = strings.Index(content, leftIndicator)
			rIndex = strings.Index(content, rightIndicator)
			if lIndex > 0 || rIndex > 0 || lIndex < rIndex {
				logData.module = strings.TrimSpace(content[lIndex+1 : rIndex])
			}
		}
	}

	lIndex = strings.Index(input, startToken)
	rIndex = strings.Index(input, endToken)
	if lIndex < 0 || rIndex < 0 {
		return nil, errors.New("line does not contains benchmark data")
	}
	var str = strings.TrimSpace(input[lIndex+startTokenLen : rIndex])
	var items = strings.Split(str, separator)

	var kvs = make([]KvPair, 0)
	for _, si := range items {
		var subStrings = strings.Split(si, ":")
		if len(subStrings) != 2 {
			continue
		}
		var val, err = strconv.Atoi(subStrings[1])
		if err != nil {
			continue
		}
		kvs = append(kvs, KvPair{subStrings[0], val})
	}
	logData.keyValues = kvs
	return &logData, nil
}
