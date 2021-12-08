package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	task01()
	//task02()

}

func task01() {
	file, err := os.Open("./input08.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var count int
	//Initial filling
	for scanner.Scan() {
		line := scanner.Text()
		str := strings.Split(line, " | ")

		//upattern := strings.Split(str[0], " ")
		output := strings.Split(str[1], " ")

		for _, str := range output {
			if len(str) == 2 || len(str) == 4 || len(str) == 3 || len(str) == 7 {
				count++
			}
		}

	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func task02() {

}
