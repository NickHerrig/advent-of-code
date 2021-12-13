package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//subtract string2 from string 1
func sub(str1, str2 string) string {
	result := str1
	for _, letter := range str2 {
		result = strings.ReplaceAll(result, string(letter), "")
	}
	return result

}

func generateSegmentMap(signals []string) []string {
	sm := map[string]string{}
	for _, signal := range signals {
		switch len(signal) {
		case 2:
			sm["1"] = signal
		case 4:
			sm["4"] = signal
		case 3:
			sm["7"] = signal
		case 7:
			sm["8"] = signal

		}
	}

	segments := make([]string, 7)
	segments[0] = sub(sm["7"], sm["1"])
	fmt.Println(segments)

	return segments
}

func decodeOutput(output string) string {
	return output
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

	for _, signal := range signals {
		sliceSignals := strings.Split(signal, " ")
		segmentSlice := generateSegmentMap(sliceSignals)
		fmt.Println(segmentSlice)
	}

}
