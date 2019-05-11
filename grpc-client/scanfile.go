package main

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

func scanFile(updfile io.Reader) (map[string][]map[int][]string, int, error) {

	data := make(map[string][]map[int][]string)
	var code int

	csvr := csv.NewReader(updfile)

	// skip header: first file line
	_, err := csvr.Read()
	if err != nil {
		log.Printf("Error reading header: %v\n", err)
	} // .if

	// skip header first file line
	csvr.Read()
	if err != nil {
		log.Printf("Error reading header: %v\n", err)
		return data, code, err
	} // .if

	// read lines:
	rowNo := 1
	for {

		line, err := csvr.Read()
		if err != nil {
			if err == io.EOF { // end of file
				break
			} // if
			log.Printf("Error reading one of the lines: %v\n", err)
			return data, code, err
		} // .if

		row := map[int][]string{rowNo: line}

		data[line[1]] = append(data[line[1]], row)

		code, _ = strconv.Atoi(line[0])

		rowNo++
	} // .for

	return data, code, nil
} // .scanFile
