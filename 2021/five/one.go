package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type line struct {
	Points [][]int
}

func (l *line) FindAllPoints() {
	x1 := l.Points[0][0]
	x2 := l.Points[1][0]
	y1 := l.Points[0][1]
	y2 := l.Points[1][1]

	vertical := x1-x2 == 0
	horizontal := y1-y2 == 0

	switch {
	case horizontal:
		min := int(math.Min(float64(x1), float64(x2)))
		max := int(math.Max(float64(x1), float64(x2)))
		for i := min + 1; i < max; i++ {
			point := []int{i, y1}
			l.Points = append(l.Points, point)
		}

	case vertical:
		min := int(math.Min(float64(y1), float64(y2)))
		max := int(math.Max(float64(y1), float64(y2)))
		for i := min + 1; i < max; i++ {
			point := []int{x1, i}
			l.Points = append(l.Points, point)
		}
	default:

	}

}

type grid struct {
	Size int
	Grid [][]int
}

func (g *grid) Plot(line *line) {

	var x, y int
	for _, p := range line.Points {
		x, y = p[0], p[1]
		g.Grid[x][y]++
	}

}

func (g *grid) Danger() int {
	var c int

	for _, row := range g.Grid {
		for _, p := range row {
			if p > 1 {
				c++
			}
		}
	}

	return c
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
	lines := []*line{}
	for scanner.Scan() {
		l, err := newLine(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		l.FindAllPoints()
		lines = append(lines, l)
	}

	// loop through lines, and plot on graph
	grid := newGrid(1000)
	for _, l := range lines {
		grid.Plot(l)
	}
	fmt.Println(grid.Danger())
}

func newGrid(size int) *grid {

	a := make([][]int, size)
	for i := range a {
		a[i] = make([]int, size)
	}

	g := &grid{
		Size: size,
		Grid: a,
	}
	return g
}

func newLine(input string) (*line, error) {

	s := strings.Split(input, " -> ")

	points := [][]int{}
	for _, val := range s {
		p := strings.Split(val, ",")
		point := []int{}
		for _, x := range p {
			i, err := strconv.Atoi(x)
			if err != nil {
				return nil, err
			}
			point = append(point, i)
		}
		points = append(points, point)
	}

	l := &line{
		Points: points,
	}

	return l, nil

}
