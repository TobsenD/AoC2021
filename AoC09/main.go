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

	task01()
	//task02()

}

func task01() {
	file, err := os.Open("./input09.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	//Prepare Data
	matrix := make([][]int, 0)
	for scanner.Scan() {
		row := make([]int, 0)
		line := scanner.Text()
		slc := strings.Split(line, "")
		for x := range slc {
			row = append(row, convertInt(slc[x]))
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//Work with data
	mHeight := len(matrix)
	mWidth := len(matrix[0])

	sum := 0

	for y, row := range matrix {
		for x, col := range row {
			left, up, right, down := x-1, y-1, x+1, y+1

			isLowPoint := true
			if left >= 0 {
				isLowPoint = col < matrix[y][left] && isLowPoint
			}

			if up >= 0 {
				isLowPoint = col < matrix[up][x] && isLowPoint
			}

			if right < mWidth {
				isLowPoint = col < matrix[y][right] && isLowPoint
			}

			if down < mHeight {
				isLowPoint = col < matrix[down][x] && isLowPoint
			}

			if isLowPoint {
				sum += col + 1
			}
		}
	}

	fmt.Println(sum)

}

func task02() {

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
