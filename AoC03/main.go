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

func comparebins(line string, compare map[int]*binarycomp) {
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

func keepInList(list []string, postiton int, value byte) []string {

	var newList []string
	for _, item := range list {
		if item[postiton] == value {
			newList = append(newList, item)
		}
	}

	return newList
}

func main() {

	//task01()
	task02()

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

		comparebins(line, compare)

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
	file, err := os.Open("./input03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var oxy []string
	var co2 []string

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		oxy = append(oxy, line)
		co2 = append(co2, line)
	}

	var oxyBit int
	for ok := true; ok; ok = len(oxy) > 1 {
		var compareOxy = map[int]*binarycomp{}
		for _, value := range oxy {
			comparebins(value, compareOxy)
		}

		if compareOxy[oxyBit].one > compareOxy[oxyBit].zero {
			oxy = keepInList(oxy, oxyBit, '1')
		} else if compareOxy[oxyBit].one == compareOxy[oxyBit].zero {
			oxy = keepInList(oxy, oxyBit, '1')
		} else {
			oxy = keepInList(oxy, oxyBit, '0')
		}

		oxyBit++

	}

	var co2Bit int
	for ok := true; ok; ok = len(co2) > 1 {
		var compareCo2 = map[int]*binarycomp{}
		for _, value := range co2 {
			comparebins(value, compareCo2)
		}

		if compareCo2[co2Bit].one > compareCo2[co2Bit].zero {
			co2 = keepInList(co2, co2Bit, '0')
		} else if compareCo2[co2Bit].one == compareCo2[co2Bit].zero {
			co2 = keepInList(co2, co2Bit, '0')
		} else {
			co2 = keepInList(co2, co2Bit, '1')
		}

		co2Bit++

	}

	if oxygen, err := strconv.ParseInt(oxy[0], 2, 64); err != nil {
		fmt.Println(err)
	} else {
		if coscrub, err := strconv.ParseInt(co2[0], 2, 64); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(oxygen * coscrub)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
