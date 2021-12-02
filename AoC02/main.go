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
	file, err := os.Open("./input02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var hor int
	var vert int

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, " ")

		switch {
		case split[0] == "forward":
			hor += convertInt(split[1])
		case split[0] == "up":
			vert -= convertInt(split[1])
		case split[0] == "down":
			vert += convertInt(split[1])
		}

	}

	fmt.Println(vert * hor)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func task02() {
	file, err := os.Open("./input02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var hor int
	var vert int
	var aim int

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, " ")

		switch {
		case split[0] == "forward":
			hor += convertInt(split[1])
			vert += aim * convertInt(split[1])
		case split[0] == "up":
			aim -= convertInt(split[1])
		case split[0] == "down":
			aim += convertInt(split[1])
		}

	}

	fmt.Println(vert * hor)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
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
