package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (b *board) update(ball string) error {
	return nil
}

func (b *board) bingo() bool {
	return false
}

func (b *board) sum() int {
	return 1000
}

type board struct {
	Numbers [][]string
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

	// read first line, split on a comma, and load bingo cage, close chanel for ranging
	line := scanner.Text()
	numbers := strings.Split(line, ",")
	c := newCage(numbers)
	close(c)

	// read lines into slice, ignoring blank lines
	lines := []string{}
	for scanner.Scan() {
		if len(scanner.Text()) > 1 {
			lines = append(lines, scanner.Text())
		}
	}

	// Initalize a slice of boards
	boards := []*board{}
	length := len(lines)
	chunk := 5
	for i := chunk; i < length; i += chunk {
		b := newBoard(lines[i-chunk : i])
		boards = append(boards, b)
	}

	// Main loop through balls, and update board
	for ball := range c {
		for _, board := range boards {
			board.update(ball)
			if sum := board.sum(); board.bingo() {
				fmt.Println("Winner!")
				b, err := strconv.Atoi(ball)
				if err != nil {
					fmt.Println("issue convering ball to int")
				}
				answer := sum * b
				fmt.Println(answer)
			}
		}

	}

}

func newBoard(lines []string) *board {

	s := [][]string{}
	for _, line := range lines {
		vals := strings.Split(line, " ")
		s = append(s, vals)
	}

	return &board{Numbers: s}

}

func newCage(numbers []string) chan string {
	buffer := len(numbers)
	c := make(chan string, buffer)
	for _, num := range numbers {
		c <- num
	}

	return c
}
