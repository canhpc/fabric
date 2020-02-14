package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"os"
)

type Printer interface {
	Print(records map[string]Record)
}

type ConsolePrinter struct {
}

func (p ConsolePrinter) Print(records map[string]Record) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Key", "Number of sample", "Mean", "Min", "Max"})
	for key, record := range records {
		t.AppendRow([]interface{}{key, record.nOfItems, record.min, record.max, fmt.Sprintf("%.2f", record.mean)})
	}
	t.Render()
}
