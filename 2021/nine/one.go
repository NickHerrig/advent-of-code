package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type heightMap struct {
	Height int
	Width  int
	Grid   [][]int
	Risk   int
}

func (hm *heightMap) GetAdjacent(x, y int) []int {

	points := []int{}
	up, err := hm.GetPoint(x, y-1)
	if err == nil {
		points = append(points, up)
	}

	down, err := hm.GetPoint(x, y+1)
	if err == nil {
		points = append(points, down)
	}

	right, err := hm.GetPoint(x+1, y)
	if err == nil {
		points = append(points, right)
	}

	left, err := hm.GetPoint(x-1, y)
	if err == nil {
		points = append(points, left)
	}

	return points
}

func (hm *heightMap) GetPoint(x, y int) (int, error) {

	if x < hm.Width && x >= 0 && y < hm.Height && y >= 0 {
		return hm.Grid[y][x], nil
	} else {
		err := errors.New("Point not on height map")
		return 0, err
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
	//	signals := []string{}
	outputs := [][]int{}
	for scanner.Scan() {
		numbers := []int{}
		row := strings.Split(scanner.Text(), "")
		for _, num := range row {
			numi, _ := strconv.Atoi(num)
			numbers = append(numbers, numi)
		}
		outputs = append(outputs, numbers)
	}

	hm := newHeightMap(outputs)
	for y, row := range outputs {
		for x, val := range row {
			adj := hm.GetAdjacent(x, y)
			low := checkLow(val, adj)
			if low {
				hm.Risk += val + 1
			}

		}
	}
	fmt.Println(hm.Risk)
}

func checkLow(val int, adj []int) bool {
	for _, a := range adj {
		if val >= a {
			return false
		}
	}
	return true
}

func newHeightMap(grid [][]int) *heightMap {
	height := len(grid)
	width := len(grid[0])

	m := &heightMap{
		Height: height,
		Width:  width,
		Grid:   grid,
	}

	return m
}
