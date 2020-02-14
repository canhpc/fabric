package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	peerLog = "/home/canhpc/workspace/text/peer1org1.log"
)

func main() {
	//open logfile
	file, err := os.Open(peerLog)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//store all log data in a dictionary(map)
	// module1 |  key1 [value_1, value_2, .. value_n]
	//         |\ key2 [value_1, value_2, .. value_n]
	//          \ key2 [value_1, value_2, .. value_n]
	//
	// module2 |  key1 [value_1, value_2, .. value_n]
	//         |\ key2 [value_1, value_2, .. value_n]
	//          \ key2 [value_1, value_2, .. value_n]
	var dict = make(map[string]map[string][]int)

	//read file line by line and pare the data.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var logData, err = Parse(scanner.Text())
		if err != nil {
			continue
		}

		// put data into global dictionary
		//if key is not exits then create is
		if _, isKeyExist := dict[logData.module]; !isKeyExist {
			dict[logData.module] = make(map[string][]int)
		}
		var temp = dict[logData.module]
		//add data to global map
		for _, di := range logData.keyValues {
			if _, ok := temp[di.key]; ok {
				temp[di.key] = append(temp[di.key], di.value)
			} else {
				temp[di.key] = []int{di.value}
			}
		}
	}

	//print result for each module
	var printer ConsolePrinter
	for module, records := range dict {
		fmt.Println("Module: ", module)
		//analyze log data
		analysisRecord := make(map[string]Record)
		for key, val := range records {
			var result, err = Analyze(val)
			if err != nil {
				continue
			}
			analysisRecord[key] = *result
		}
		printer.Print(analysisRecord)
	}
}
