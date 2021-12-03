package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type binarycomp struct {
	zero int
	one  int
}

func newBinaryComp() *binarycomp {
	bc := binarycomp{one: 0, zero: 0}
	return &bc
}

func main() {

	task01()
	//task02()

}

func task01() {
	file, err := os.Open("./input03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var compare = map[int]*binarycomp{}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			if compare[i] == nil {
				compare[i] = newBinaryComp()
			}

			switch {
			case line[i] == '1':
				compare[i].one += 1
			case line[i] == '0':
				compare[i].zero += 1
			}
		}

	}

	var keys []int
	for k := range compare {
		keys = append(keys, k)
	}

	var gamma string = ""
	var epsilon string = ""

	sort.Ints(keys)
	for _, k := range keys {

		if compare[k].one > compare[k].zero {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}

	if gnumber, err := strconv.ParseInt(gamma, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		if cnumber, err := strconv.ParseInt(epsilon, 2, 64); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(gnumber * cnumber)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func task02() {

}
