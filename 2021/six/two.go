package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type school struct {
	Fish map[int]int
}

func (s *school) Spawn() {
	for fish, days := range s.Fish {
		if days > 0 {
			s.Fish[fish]--
		} else {
			s.Fish[fish] = 6
			s.Fish[len(s.Fish)+1] = 8
		}
	}
}

func (s *school) Count() int {
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

	for i := 0; i < 256; i++ {
		school.Spawn()
		fmt.Println("Day:", i, "Count:", school.Count())
	}

}

func newSchool(fish []string) *school {

	fishies := map[int]int{}
	for i, f := range fish {
		fi, err := strconv.Atoi(f)
		if err != nil {
			fmt.Println("issue convering fish to int")
		}
		fishies[i] = fi
	}

	return &school{Fish: fishies}

}
