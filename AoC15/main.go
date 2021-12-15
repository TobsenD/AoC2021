package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/RyanCarrier/dijkstra"
)

func main() {

	task()
}

type coords struct {
	x int
	y int
}

func task() {
	file, err := os.Open("./input15.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	//Prepare Data
	matrix := make([][]int64, 0)
	for scanner.Scan() {
		row := make([]int64, 0)
		line := scanner.Text()
		slc := strings.Split(line, "")
		for x := range slc {
			row = append(row, int64(convertInt(slc[x])))
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//Work with data
	mHeight := len(matrix)
	mWidth := len(matrix[0])

	//create node map
	var count int
	nodeMap := make(map[coords]int)
	graph := dijkstra.NewGraph()
	for y, row := range matrix {
		for x := range row {
			nodeMap[coords{x, y}] = count
			graph.AddVertex(count)
			count++
		}
	}

	//creating arcs
	for y, row := range matrix {
		for x := range row {
			left, up, right, down := x-1, y-1, x+1, y+1

			if left >= 0 {
				graph.AddArc(nodeMap[coords{x, y}], nodeMap[coords{left, y}], matrix[y][left])
			}

			if up >= 0 {
				graph.AddArc(nodeMap[coords{x, y}], nodeMap[coords{x, up}], matrix[up][x])
			}

			if right < mWidth {
				graph.AddArc(nodeMap[coords{x, y}], nodeMap[coords{right, y}], matrix[y][right])
			}

			if down < mHeight {
				graph.AddArc(nodeMap[coords{x, y}], nodeMap[coords{x, down}], matrix[down][x])
			}
		}
	}

	best, err := graph.Shortest(0, count-1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shortest distance ", best.Distance, " following path ", best.Path)

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
