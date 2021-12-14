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
	Lows   []int
}

func (hm *heightMap) GetAdjacent(x, y int) ([]int, error) {
	switch {
	case x == 0 && y == 0:
		down, err := hm.GetPoint(0, 1)
		right, err := hm.GetPoint(1, 0)
		if err != nil {
			return nil, err
		}
		return []int{down, right}, nil
	}

	err := errors.New("Counln't get adjacent points")
	return nil, err

}

func (hm *heightMap) GetPoint(x, y int) (int, error) {

	if x < hm.Width && y < hm.Height {
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
	pts, err := hm.GetAdjacent(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pts)

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
