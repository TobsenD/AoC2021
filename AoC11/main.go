package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	//task01()
	task02()

}

func task01() {
	file, err := os.Open("./input11.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	//Prepare Data
	grid := make([][]int, 0)
	for scanner.Scan() {
		row := make([]int, 0)
		line := scanner.Text()
		slc := strings.Split(line, "")
		for x := range slc {
			row = append(row, convertInt(slc[x]))
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var flashCount int
	steps := 100

	for i := 0; i < steps; i++ {
		//First every energylevel increase by one
		for y, row := range grid {
			for x := range row {
				grid[y][x]++
			}
		}

		for y, row := range grid {
			for x, col := range row {
				if col > 9 {
					flashCount += flash(grid, x, y)
				}
			}
		}

	}

	fmt.Println(flashCount)
}

func task02() {
	file, err := os.Open("./input11.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	//Prepare Data
	grid := make([][]int, 0)
	for scanner.Scan() {
		row := make([]int, 0)
		line := scanner.Text()
		slc := strings.Split(line, "")
		for x := range slc {
			row = append(row, convertInt(slc[x]))
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var flashCount int

	allFlashing := false
	for i := 0; !allFlashing; i++ {
		//First every energylevel increase by one
		for y, row := range grid {
			for x := range row {
				grid[y][x]++
			}
		}

		for y, row := range grid {
			for x, col := range row {
				if col > 9 {
					flashCount += flash(grid, x, y)
				}
			}
		}

		allFlashing = true
		for _, row := range grid {
			for _, col := range row {
				if col != 0 {
					allFlashing = false
				}
			}
		}

		if allFlashing {
			fmt.Println(grid)
			fmt.Println("STEP: ", i+1)
		}

	}
}

func flash(grid [][]int, x int, y int) int {

	var flashCount int

	flashCount++
	grid[y][x] = 0

	//Work with data
	mHeight := len(grid)
	mWidth := len(grid[0])
	left, up, right, down := x-1, y-1, x+1, y+1

	if left >= 0 && up >= 0 {
		if grid[up][left] != 0 {
			grid[up][left]++
			if grid[up][left] > 9 {
				flashCount += flash(grid, left, up)
			}
		}
	}

	if up >= 0 {
		if grid[up][x] != 0 {
			grid[up][x]++
			if grid[up][x] > 9 {
				flashCount += flash(grid, x, up)
			}
		}
	}

	if right < mWidth && up >= 0 {
		if grid[up][right] != 0 {
			grid[up][right]++
			if grid[up][right] > 9 {
				flashCount += flash(grid, right, up)
			}
		}
	}

	if right < mWidth {
		if grid[y][right] != 0 {
			grid[y][right]++
			if grid[y][right] > 9 {
				flashCount += flash(grid, right, y)
			}
		}
	}

	if right < mWidth && down < mHeight {
		if grid[down][right] != 0 {
			grid[down][right]++
			if grid[down][right] > 9 {
				flashCount += flash(grid, right, down)
			}
		}
	}

	if down < mHeight {
		if grid[down][x] != 0 {
			grid[down][x]++
			if grid[down][x] > 9 {
				flashCount += flash(grid, x, down)
			}
		}
	}

	if left >= 0 && down < mHeight {
		if grid[down][left] != 0 {
			grid[down][left]++
			if grid[down][left] > 9 {
				flashCount += flash(grid, left, down)
			}
		}
	}

	if left >= 0 {
		if grid[y][left] != 0 {
			grid[y][left]++
			if grid[y][left] > 9 {
				flashCount += flash(grid, left, y)
			}
		}
	}

	return flashCount

}

func convertInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}
