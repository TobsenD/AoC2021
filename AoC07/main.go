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
	file, err := os.Open("./input07.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var crabSubs []int

	//Initial filling
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), ",")
		for _, str := range strs {
			crabSubs = append(crabSubs, convertInt(str))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	minCrab, maxCrab := minMax(crabSubs)

	fuel := calcCrabFuel(crabSubs, minCrab, maxCrab)

	fmt.Println(fuel)
}

func task02() {
	file, err := os.Open("./input07.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var crabSubs []int

	//Initial filling
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), ",")
		for _, str := range strs {
			crabSubs = append(crabSubs, convertInt(str))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	minCrab, maxCrab := minMax(crabSubs)

	fuel := calcCrabEngineeredFuel(crabSubs, minCrab, maxCrab)

	fmt.Println(fuel)
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

func minMax(array []int) (min, max int) {
	for _, value := range array {
		if value < min {
			min = value
		} else if value > max {
			max = value
		}
	}
	return min, max
}

func calcCrabFuel(array []int, min, max int) int {
	var cheapestFuel int
	for minMaxValue := min; minMaxValue <= max; minMaxValue++ {
		var tempFuelCost int
		for _, value := range array {
			if value < minMaxValue {
				tempFuelCost += minMaxValue - value
			} else if value > minMaxValue {
				tempFuelCost += value - minMaxValue
			} else if value == minMaxValue {
				tempFuelCost += 0
			}
		}
		if tempFuelCost < cheapestFuel || cheapestFuel == 0 {
			cheapestFuel = tempFuelCost
		}
	}
	return cheapestFuel
}

func calcCrabEngineeredFuel(array []int, min, max int) int {
	var cheapestFuel int
	for minMaxValue := min; minMaxValue <= max; minMaxValue++ {
		var tempFuelCost int
		for _, value := range array {
			if value < minMaxValue {
				steps := minMaxValue - value
				for x := 0; x < steps; x++ {
					tempFuelCost += x + 1
				}
			} else if value > minMaxValue {
				steps := value - minMaxValue
				for x := 0; x < steps; x++ {
					tempFuelCost += x + 1
				}
			} else if value == minMaxValue {
				tempFuelCost += 0
			}
		}
		if tempFuelCost < cheapestFuel || cheapestFuel == 0 {
			cheapestFuel = tempFuelCost
		}
	}
	return cheapestFuel
}
