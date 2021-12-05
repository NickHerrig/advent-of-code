package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
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

	increasedCount := 0
	for i, val := range input {

		if i == 0 {
			continue
		}

		prev := input[i-1]

		if val > prev {
			increasedCount++
		}
	}
	fmt.Println(increasedCount)

}
