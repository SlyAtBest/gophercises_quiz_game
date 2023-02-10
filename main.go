package main

import (
	"flag"
	"fmt"
	"quiz/input"
)

func main() {
	var fileFlag = flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	reader := input.NewReader(*fileFlag)
	_, err := reader.ReadData()

	if err != nil {
		fmt.Println("Fatal error occurred: ", err)
		return
	}
}
