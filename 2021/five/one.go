package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	Rise   int
	Run    int
	Points []int
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

	// read lines into slice, ignoring blank lines
	lines := []*Line{}
	for scanner.Scan() {
		l, err := newLine(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, l)
		fmt.Println(l.Points)
	}

}

func newLine(line string) (*Line, error) {
	l := &Line{}
	s := strings.Split(line, " -> ")
	for _, val := range s {
		p := strings.Split(val, ",")
		for _, x := range p {
			i, err := strconv.Atoi(x)
			if err != nil {
				return nil, err
			}
			l.Points = append(l.Points, i)
		}
	}

	return l, nil

}
