package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type school struct {
	Fish []int
}

func (s *school) Spawn() {
	for i, fish := range s.Fish {
		if fish > 0 {
			s.Fish[i]--
		} else {
			s.Fish[i] = 6
			s.Fish = append(s.Fish, 8)
		}

	}
}

func (s *school) Count() int64 {
	return len(s.Fish)
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

	for i := 0; i < 80; i++ {
		school.Spawn()
	}

	fmt.Println(school.Count())

}

func newSchool(fish []string) *school {

	fishies := []int{}
	for _, f := range fish {
		i, err := strconv.Atoi(f)
		if err != nil {
			fmt.Println("issue convering fish to int")
		}
		fishies = append(fishies, i)
	}

	return &school{Fish: fishies}

}
