package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Conversion(filepath string) {

	var err error
	var lines [][]string

	csvFile, err := os.Open(filepath)

	if err != nil {
		BadPath("Failed to open File!")
		return
	}

	csv := csv.NewReader(csvFile)

	filestring, _ := ioutil.ReadFile(filepath)

	if strings.Contains(string(filestring), ";") {

		csv.Comma = ';'

	}

	if strings.Contains(string(filestring), "\t") {
		csv.Comma = '\t'
	}

	lines, err = csv.ReadAll()

	if err != nil {
		BadPath(fmt.Sprintf("%v", err))
	}

	keyscan := true

	var keys []string

	var parts []string

	var data []string

	for _, line := range lines {

		if keyscan == true {
			for _, key := range line {
				keys = append(keys, key)
			}
			keyscan = false
			continue
		}

		for x, y := range keys {

			parts = append(parts, `"`+y+`":"`+line[x]+`"`)

		}

		part := strings.Join(parts, ",")

		data = append(data, "{"+part+"}")

	}

	jsondata := "[" + strings.Join(data, ",") + "]"

	filename := SaveDialog()

	err = ioutil.WriteFile(filename, []byte(jsondata), 0777)

	if err != nil {
		BadPath("Failed to save JSON-File:\n\n" + fmt.Sprintf("%v", err))

	}

}
