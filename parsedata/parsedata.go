package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var fileName = "scriptDATA.txt"

// President ...
type President struct {
	Name            string
	AgeInauguration float64
}

// Month
const (
	_ = iota
	Jan
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)

func main() {
	data, _ := ioutil.ReadFile(fileName)
	dataLines := strings.Split(string(data), "\n")
	for i := range dataLines {
		dataLines[i] = strings.SplitAfterN(dataLines[i], "   ", 2)[1]
		dataLines[i] = strings.TrimLeft(dataLines[i], " ")
		dataLines[i] = strings.Split(dataLines[i], ", ")[2]
		fmt.Println(strings.Split(dataLines[i], " ")[0])
		// strings.Split(dataLines[i], ", ")[1]
	}
}
