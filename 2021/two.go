package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(array []int) int {
	result := 0
	for _, val := range array {
		result += val
	}
	return result
}

func main() {
	file, err := os.Open("input-one.txt")
	if err != nil {
		fmt.Println("issue opening input file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []int
	for scanner.Scan() {
		t := scanner.Text()
		i, err := strconv.Atoi(t)
		if err != nil {
			fmt.Println("issue convering string to uint64")
			os.Exit(1)
		}
		input = append(input, i)
	}

	l := len(input) % 3
	r := len(input) - l
	var data []int
	for i, val := range input[:r] {
		bucket := []int{val, input[i+1], input[i+2]}
		s := sum(bucket)
		data = append(data, s)
	}

	increasedCount := 0
	for i, val := range data {

		if i == 0 {
			continue
		}

		prev := data[i-1]

		if val > prev {
			increasedCount++
		}
	}

	fmt.Println(increasedCount)

}
