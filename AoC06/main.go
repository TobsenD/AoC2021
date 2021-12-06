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
	file, err := os.Open("./input06.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var fishTimers []int
	days := 80

	//Initial filling
	for scanner.Scan() {
		//Reading first Line with Bingo numbers.
		strs := strings.Split(scanner.Text(), ",")
		for _, str := range strs {
			fishTimers = append(fishTimers, convertInt(str))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < days; i++ {
		//fmt.Println("After %s days: %s", i, fishTimers)
		spawnFish := 0
		for j := 0; j < len(fishTimers); j++ {
			if fishTimers[j] == 0 {
				spawnFish++
				fishTimers[j] = 6
			} else {
				fishTimers[j]--
			}
		}

		if spawnFish > 0 {
			for j := 0; j < spawnFish; j++ {
				fishTimers = append(fishTimers, 8)
			}
		}
	}

	fmt.Println(len(fishTimers))

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
