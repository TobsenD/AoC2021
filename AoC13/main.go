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

	task()

}

type coords struct {
	x int
	y int
}

func task() {
	file, err := os.Open("./input13.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	paper := make(map[coords]bool)
	var fold []string
	scanner := bufio.NewScanner(file)
	//Build Grid and Fold Commands
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "fold") {
			fold = append(fold, strings.Trim(line, "fold along "))
		} else if line != "" {
			str := strings.Split(line, ",")
			paper[coords{convertInt(str[0]), convertInt(str[1])}] = true
		}
	}

	for _, ele := range fold {
		command := strings.Split(ele, "=")

		if command[0] == "y" {
			paper = foldUp(paper, convertInt(command[1]))
		} else {
			paper = foldLeft(paper, convertInt(command[1]))
		}
	}

	printPaper(paper)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func foldUp(paper map[coords]bool, foldline int) map[coords]bool {

	newPaper := make(map[coords]bool)

	for nY := 0; nY < foldline; nY++ {
		for position := range paper {
			if position.y == nY || position.y == foldline+foldline-nY {
				newPaper[coords{position.x, nY}] = true
			}
		}
	}

	if (!newPaper[coords{0, foldline - 1}]) {
		newPaper[coords{0, foldline - 1}] = false
	}

	return newPaper
}

func foldLeft(paper map[coords]bool, foldline int) map[coords]bool {

	newPaper := make(map[coords]bool)

	for nX := 0; nX < foldline; nX++ {
		for position := range paper {
			if position.x == nX || position.x == foldline+foldline-nX {
				newPaper[coords{nX, position.y}] = true
			}
		}
	}

	if (!newPaper[coords{foldline - 1, 0}]) {
		newPaper[coords{foldline - 1, 0}] = false
	}

	return newPaper
}

func printPaper(paper map[coords]bool) {

	var countDots int

	var maxX int
	var maxY int
	for k := range paper {

		if maxX <= k.x {
			maxX = k.x
		}
		if maxY <= k.y {
			maxY = k.y
		}

	}

	for row := 0; row <= maxY; row++ {
		for col := 0; col <= maxX; col++ {
			if (paper[coords{col, row}]) {
				fmt.Print("#")
				countDots++
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println(countDots)
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
