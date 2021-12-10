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
	file, err := os.Open("./input10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	opening := map[string]bool{
		"(": true,
		"<": true,
		"[": true,
		"{": true,
	}
	closing := map[string]bool{
		")": true,
		">": true,
		"]": true,
		"}": true,
	}
	opposing := map[string]string{
		")": "(",
		">": "<",
		"]": "[",
		"}": "{",
	}
	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	var sum int

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "")
		var stack []string

		for _, v := range tokens {
			if opening[v] {
				stack = append(stack, v)
			}
			if closing[v] {
				tail := stack[len(stack)-1]
				if tail != opposing[v] {
					sum += points[v]
					break
				}
				stack = stack[:len(stack)-1]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func task02() {

}
