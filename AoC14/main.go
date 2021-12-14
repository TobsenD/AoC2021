package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {

	task()

}

func task() {
	file, err := os.Open("./input14.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var origin []string
	rules := make(map[string][]string)
	ruleGroup := make(map[string]int)

	//Build Grid and Fold Commands
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, " -> ") {
			tmpstr := strings.Split(line, " -> ")
			strs := strings.Split(tmpstr[0], "")

			rule := []string{strs[0] + tmpstr[1], tmpstr[1] + strs[1]}

			rules[tmpstr[0]] = rule

		} else if line != "" {
			origin = strings.Split(line, "")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(origin)-1; i++ {
		rule := origin[i] + origin[i+1]
		ruleGroup[rule] = 1
	}

	fmt.Println(ruleGroup)
	fmt.Println(rules)

	steps := 10

	for i := 0; i < steps; i++ {
		ruleGroup = insertPolymere(ruleGroup, rules)
	}

	fmt.Println(ruleGroup)

	count(ruleGroup)

}

func insertPolymere(ruleGroup map[string]int, rules map[string][]string) map[string]int {
	newRuleGroup := make(map[string]int)

	for k, v := range ruleGroup {

		for i := 0; i < v; i++ {
			rule1 := rules[k][0]
			rule2 := rules[k][1]

			//fmt.Println("For %d, setting %s and %s", k, rule1, rule2)

			newRuleGroup[rule1]++
			newRuleGroup[rule2]++
		}
	}

	return newRuleGroup
}

func count(ruleGroup map[string]int) {
	stringCount := make(map[string]int)
	for k, v := range ruleGroup {
		strs := strings.Split(k, "")
		for _, ele := range strs {
			stringCount[ele] = stringCount[ele] + v
		}

	}

	var min float64 = 9999999999999
	var max float64

	for _, v := range stringCount {
		f := float64(v) / float64(2)
		if math.Round(f) >= max {
			max = math.Round(f)
		}
		if math.Round(f) <= min {
			min = math.Round(f)
		}
	}

	fmt.Println("Task1: ", max-min)

}
