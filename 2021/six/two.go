package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Spawn(school []int64) []int64 {

	newFish := school[0]

	// decrement fish by one day
	for i := 0; i < 8; i++ {
		school[i] = school[i+1]
	}

	//new fish
	school[8] = newFish

	//reset fish
	school[6] += newFish

	return school
}

func Count(school []int64) int64 {

	var count int64
	for _, val := range school {
		count += val
	}

	return count

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
	school := newSchool(input)

	for i := 0; i < 256; i++ {
		school = Spawn(school)
	}
	fmt.Println(Count(school))

}

func newSchool(fish []string) []int64 {

	school := make([]int64, 9)
	for _, f := range fish {
		fi, err := strconv.Atoi(f)
		if err != nil {
			fmt.Println("issue convering fish to int")
		}
		school[fi]++
	}

	return school

}
