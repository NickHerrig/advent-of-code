package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type board struct {
	Numbers [][]int
}

func (b *board) update(ball int) error {
	for x, row := range b.Numbers {
		for y, val := range row {
			if ball == val {
				b.Numbers[x][y] = 0
			}
		}
	}

	return nil
}
func (b *board) check(line []int) bool {
	var sum int
	for _, val := range line {
		sum += val
	}
	if sum == 0 {
		return true
	}

	return false
}

func (b *board) bingo() bool {
	// check rows first...
	for _, row := range b.Numbers {
		if b.check(row) {
			return true
		}
	}

	for y := 0; y < 5; y++ {
		column := []int{}
		for x, _ := range b.Numbers {
			column = append(column, b.Numbers[x][y])
		}
		if b.check(column) {
			return true
		}
	}

	return false
}

func (b *board) sum() int {

	var sum int
	for _, row := range b.Numbers {
		for _, val := range row {
			sum += val
		}
	}

	return sum
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
				fmt.Println("Found Winner!")
				answer := sum * ball
				fmt.Println(answer)
				os.Exit(0)
			}
		}

	}

}

func newBoard(lines []string) *board {

	b := [][]int{}

	for _, line := range lines {
		vals := strings.Split(line, " ")
		n := []int{}
		for _, val := range vals {
			valInt, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			n = append(n, valInt)
		}
		b = append(b, n)

	}

	return &board{Numbers: b}

}

func newCage(numbers []string) chan int {
	buffer := len(numbers)
	c := make(chan int, buffer)
	for _, num := range numbers {
		b, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("issue convering ball to int")
		}
		c <- b
	}

	return c
}
