package main

import (
	"bufio"
	"fmt"
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

//calculates fuel for one crab
func calculateFuel(i, crab int) int {
	start := min([]int{i, crab})
	end := max([]int{i, crab})

	fuelInc := 1
	fuel := 0
	for i := start; i < end; i++ {
		fuel += fuelInc
		fuelInc++
	}

	return fuel

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
	// range over each horizontal position for testing
	for i := start; i < end; i++ {

		// calculate fuel for crab moving to horizontal position
		// sum the totalFuel for moving all crabs
		var totalFuel int
		for _, crab := range positions {
			totalFuel += calculateFuel(i, crab)
		}

		sliceFuels = append(sliceFuels, totalFuel)
	}

	fmt.Println(min(sliceFuels))

}
