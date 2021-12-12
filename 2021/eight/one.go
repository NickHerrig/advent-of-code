package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	ONE   = 2
	FOUR  = 4
	SEVEN = 3
	EIGHT = 7
)

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("issue opening input file")
		os.Exit(1)
	}
	defer file.Close()

	// create a scanner
	scanner := bufio.NewScanner(file)
	//	signals := []string{}
	outputs := []string{}
	for scanner.Scan() {
		val := strings.Split(scanner.Text(), " | ")
		//	signals = append(signals, val[0])
		outputs = append(outputs, val[1])
	}

	count := 0
	for _, output := range outputs {
		sliceVals := strings.Split(output, " ")
		for _, o := range sliceVals {
			switch len(o) {
			case ONE:
				count++
			case FOUR:
				count++
			case SEVEN:
				count++
			case EIGHT:
				count++
			}

		}
	}
	fmt.Println(count)

}
