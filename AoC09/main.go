package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	//task01()
	task02()

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

	countList := make([]int, 0)

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
				count := calcBasinSize(matrix, x, y, 1, make(map[string]bool))
				countList = append(countList, count)
			}
		}
	}

	sort.Ints(countList)
	countlen := len(countList)
	fmt.Println(countList[countlen-1] * countList[countlen-2] * countList[countlen-3])
}

func calcBasinSize(matrix [][]int, x int, y int, count int, visited map[string]bool) int {
	position := fmt.Sprintf("%d,%d", y, x)
	if _, isVisited := visited[position]; isVisited {
		return count - 1
	} else {
		visited[position] = true
	}

	if matrix[y][x] == 9 {
		return count - 1
	}

	mHeight := len(matrix)
	mWidth := len(matrix[0])

	left, up, right, down := x-1, y-1, x+1, y+1

	if left >= 0 {
		count = calcBasinSize(matrix, left, y, count+1, visited)
	}

	if up >= 0 {
		count = calcBasinSize(matrix, x, up, count+1, visited)
	}

	if right < mWidth {
		count = calcBasinSize(matrix, right, y, count+1, visited)
	}

	if down < mHeight {
		count = calcBasinSize(matrix, x, down, count+1, visited)
	}

	return count
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
