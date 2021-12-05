package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	task01()
	//task02()

}

func task01() {
	file, err := os.Open("./input05.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var ventMap [1000][1000]int

	for scanner.Scan() {
		line := scanner.Text()

		array := regexp.MustCompile("[\\,\\->\\s]+").Split(line, -1)
		var ventPosition []int
		for _, value := range array {
			ventPosition = append(ventPosition, convertInt(value))
		}

		if ventPosition[0] == ventPosition[2] {
			if ventPosition[1] < ventPosition[3] {
				for i := ventPosition[1]; i <= ventPosition[3]; i++ {
					ventMap[ventPosition[0]][i]++
				}
			} else {
				for i := ventPosition[1]; i >= ventPosition[3]; i-- {
					ventMap[ventPosition[0]][i]++
				}
			}
		}
		if ventPosition[1] == ventPosition[3] {
			if ventPosition[0] < ventPosition[2] {
				for i := ventPosition[0]; i <= ventPosition[2]; i++ {
					ventMap[i][ventPosition[1]]++
				}
			} else {
				for i := ventPosition[0]; i >= ventPosition[2]; i-- {
					ventMap[i][ventPosition[1]]++
				}
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	//fmt.Println(ventMap)

	var ventCount int

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if ventMap[i][j] > 1 {
				ventCount++
			}
		}
	}

	fmt.Println(ventCount)

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
