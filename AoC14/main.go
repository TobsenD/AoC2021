package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	task()
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)

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
	ruleGroup := make(map[string]float64)

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

	steps := 40

	for i := 0; i < steps; i++ {
		ruleGroup = insertPolymere(ruleGroup, rules)
	}

	fmt.Println(ruleGroup)

	count(ruleGroup)

}

func insertPolymere(ruleGroup map[string]float64, rules map[string][]string) map[string]float64 {
	newRuleGroup := make(map[string]float64)

	for k, v := range ruleGroup {
		rule1 := rules[k][0]
		rule2 := rules[k][1]
		newRuleGroup[rule1] = newRuleGroup[rule1] + v
		newRuleGroup[rule2] = newRuleGroup[rule2] + v
	}

	return newRuleGroup
}

func count(ruleGroup map[string]float64) {
	stringCount := make(map[string]float64)
	for k, v := range ruleGroup {
		strs := strings.Split(k, "")
		for _, ele := range strs {
			stringCount[ele] = stringCount[ele] + v
		}

	}

	var min float64 = 9999999999999
	var max float64

	for k, v := range stringCount {
		f := v / 2
		if math.Round(f) >= max {
			max = math.Round(f)
		}
		if math.Round(f) <= min {
			min = math.Round(f)
		}

		fmt.Println(k, math.Round(f))
	}

	//fmt.Println("Task1: ")
	fmt.Printf("%f\n", max-min)
}
