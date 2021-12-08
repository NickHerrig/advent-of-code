package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func filter(input []string, position int, criteria string) []string {

	filtered := []string{}
	for _, val := range input {
		if string(val[position]) == criteria {
			filtered = append(filtered, val)
		}
	}
	return filtered
}

func bitSummary(input []string, position int) map[string]int64 {
	s := map[string]int64{}
	for _, val := range input {
		bit := string(val[position])
		s[bit]++
	}

	return s
}

func oxygenRate(input []string) (int64, error) {
	// Most common bit

	bitLength := len(input[0])

	for i := 0; i < bitLength; i++ {
		s := bitSummary(input, i)

		if s["1"] >= s["0"] {
			// filter input on one
			input = filter(input, i, "1")
		} else {
			// filter input filter input on two
			input = filter(input, i, "0")
		}

		if len(input) == 1 {
			return strconv.ParseInt(input[0], 2, 64)
		}

	}
	return 0, errors.New("Could not filter down to one data point")

}

func co2Rate(input []string) (int64, error) {
	// Most common bit

	bitLength := len(input[0])

	for i := 0; i < bitLength; i++ {
		s := bitSummary(input, i)

		if s["1"] >= s["0"] {
			// filter input on one
			input = filter(input, i, "0")
		} else {
			// filter input filter input on two
			input = filter(input, i, "1")
		}

		if len(input) == 1 {
			return strconv.ParseInt(input[0], 2, 64)
		}

	}
	return 0, errors.New("Could not filter down to one data point")

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

	g, err := oxygenRate(input)
	if err != nil {
		fmt.Print("error converting oxygen rate")
		fmt.Print(err)
	}

	e, err := co2Rate(input)
	if err != nil {
		fmt.Print("error converting c02 rate")
	}

	fmt.Println(g * e)

}
