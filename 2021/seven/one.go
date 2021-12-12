package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func max(slice []int) int {

	max := slice[0]
	for _, num := range slice {
		if num >= max {
			max = num
		}
	}
	return max
}

func min(slice []int) int {

	min := slice[0]
	for _, num := range slice {
		if num <= min {
			min = num
		}
	}
	return min
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
	scanner.Scan()
	input := strings.Split(scanner.Text(), ",")

	positions := []int{}
	for _, i := range input {
		num, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("Issue convering string to float")
		}

		positions = append(positions, num)

	}

	start := min(positions)
	end := max(positions)

	sliceFuels := []int{}
	for i := start; i < end; i++ {

		var totalFuel int
		for _, pos := range positions {
			totalFuel += int(math.Abs(float64(i) - float64(pos)))
		}

		sliceFuels = append(sliceFuels, totalFuel)
	}

	fmt.Println(min(sliceFuels))
	fmt.Println(max(sliceFuels))

}
