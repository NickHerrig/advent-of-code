package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
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

	open := "([{<"
	closed := ")]}>"

	pointsMap := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	openClosedMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}

	totalPoints := 0
	i := 0
	for scanner.Scan() {
		i++
		line := strings.Split(scanner.Text(), "")
		var stack Stack
		for _, del := range line {
			if strings.Contains(open, del) {
				stack.Push(del)
				continue
			}

			if strings.Contains(closed, del) {
				check, _ := stack.Pop()
				if openClosedMap[check] == del {
				} else {
					totalPoints += pointsMap[del]
					break
				}
			}
		}
	}

	fmt.Println(totalPoints)
}
