package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	//task01()
	task02()

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
		"(": ")",
		"<": ">",
		"[": "]",
		"{": "}",
	}
	points := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	var scores []int

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "")
		var stack []string
		var isValid = true
		for _, v := range tokens {
			if opening[v] {
				stack = append(stack, v)
			}
			if closing[v] {
				tail := stack[len(stack)-1]
				if tail != opposing[v] {
					isValid = false
					break
				}
				stack = stack[:len(stack)-1]
			}
		}

		var sum int
		if isValid {
			for i := range stack {
				s := stack[len(stack)-1-i]
				sum *= 5
				sum += points[opposing[s]]
			}
			scores = append(scores, sum)
		}

	}

	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
