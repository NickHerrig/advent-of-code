package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type submarine struct {
	Depth    int
	Position int
}

func (s *submarine) Move(command string, value int) {
	switch command {
	case "forward":
		s.Position += value
	case "down":
		s.Depth += value
	case "up":
		s.Depth -= value
	}
}

func (s *submarine) Answer() int {
	return s.Depth * s.Position
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

	sub := &submarine{
		Depth:    0,
		Position: 0,
	}
	for _, command := range input {
		s := strings.Split(command, " ")
		command := s[0]
		value, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println("Not able to convert string value to Int")
		}

		sub.Move(command, value)
	}

	fmt.Println(sub.Answer())

}
