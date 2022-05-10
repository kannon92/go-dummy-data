/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fptr := flag.String("input", "test.txt", "file path to read from")
	outPtr := flag.String("output", "", "Output file")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)
	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	out, err := os.Create(*outPtr)
	if err != nil {
		log.Fatal((err))
	}
	_, outErr := out.WriteString(data[0][0])
	if outErr != nil {
		log.Fatal(outErr)
	}
}
