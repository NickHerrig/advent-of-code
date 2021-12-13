package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//const (
//	ONE   = 2
//	FOUR  = 4
//	SEVEN = 3
//	EIGHT = 7
//)

func generateSegmentMap(signals []string) map[string]string {
	sm := map[string]string{
		"abcefg":  "0",
		"cf":      "1",
		"acdeg":   "2",
		"acdfg":   "3",
		"bcdf":    "4",
		"abdfg":   "5",
		"abdefg":  "6",
		"acf":     "7",
		"abcdefg": "8",
		"abcdfg":  "9",
	}

	return sm
}

func decodeOutput(segmentMap map[string]string, output string) int {
	decoded := ""
	nums := strings.Split(output, " ")
	for _, num := range nums {
		segmentMap[num] += decoded

	}

	fmt.Println(decoded)
	decodedInt, err := strconv.Atoi(decoded)
	if err != nil {
		fmt.Println("Issue parsing string to int")
	}

	return decodedInt
}

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
	signals := []string{}
	outputs := []string{}
	for scanner.Scan() {
		val := strings.Split(scanner.Text(), " | ")
		signals = append(signals, val[0])
		outputs = append(outputs, val[1])
	}

	//Main Loop for sum of outputs
	sumOutput := 0
	for i, signal := range signals {
		sliceSignals := strings.Split(signal, " ")
		segmentMap := generateSegmentMap(sliceSignals)
		outputDecoded := decodeOutput(segmentMap, outputs[i])
		sumOutput += outputDecoded
	}

	fmt.Println(sumOutput)

}
