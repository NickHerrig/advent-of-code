package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func bitSummary(input []string, position int) map[string]int64 {
	s := map[string]int64{}
	for _, val := range input {
		bit := string(val[position])
		s[bit]++
	}

	return s
}

func gammaRate(input []string) (int64, error) {
	// Most common bit

	r := ""
	bitLength := len(input[0])

	for i := 0; i < bitLength; i++ {
		s := bitSummary(input, i)

		if s["1"] > s["0"] {
			r += "1"
		} else {
			r += "0"
		}
	}

	return strconv.ParseInt(r, 2, 64)
}

func epsilonRate(input []string) (int64, error) {
	// Most common bit

	r := ""
	bitLength := len(input[0])

	for i := 0; i < bitLength; i++ {
		s := bitSummary(input, i)

		if s["1"] > s["0"] {
			r += "0"
		} else {
			r += "1"
		}
	}

	return strconv.ParseInt(r, 2, 64)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("issue opening input file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		t := scanner.Text()
		input = append(input, t)
	}

	g, err := gammaRate(input)
	if err != nil {
		fmt.Print("error converting gamma rate")
	}

	e, err := epsilonRate(input)
	if err != nil {
		fmt.Print("error converting gamma rate")
	}

	fmt.Println(g * e)

}
