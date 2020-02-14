package main

import "time"

type LogData struct {
	module    string
	timStamp  time.Time
	keyValues []KvPair
}

type KvPair struct {
	key   string
	value int
}

type Record struct {
	nOfItems int
	min      int
	max      int
	mean     float64
}
