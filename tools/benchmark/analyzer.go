package main

import "errors"

func Analyze(values []int) (*Record, error) {
	if len(values) < 1 {
		return nil, errors.New("input is empty")
	}
	var record Record
	record.min = values[0]
	record.max = values[0]
	record.nOfItems = len(values)
	var sum = 0

	for _, val := range values {
		if val < record.min {
			record.min = val
		}
		if val > record.max {
			record.max = val
		}
		sum += val
	}
	record.mean = float64(sum) / float64(record.nOfItems)
	return &record, nil
}
