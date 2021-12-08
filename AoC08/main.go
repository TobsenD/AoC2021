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

		//Teststring
		dict := buildDict(str[0])
		//fmt.Println(dict)

		output := strings.Split(str[1], " ")

		//fmt.Println(output, (getNumber(dict, output[0])*1000)+(getNumber(dict, output[1])*100)+(getNumber(dict, output[2])*10)+(getNumber(dict, output[3])))
		count = count + (getNumber(dict, output[0]) * 1000) + (getNumber(dict, output[1]) * 100) + (getNumber(dict, output[2]) * 10) + (getNumber(dict, output[3]))

	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func getNumber(dict map[int]string, str string) int {

	var num int
	for k, v := range dict {

		str1 := SortString(str)
		str2 := SortString(v)

		if str1 == str2 {
			num = k
		}
	}

	return num
}

func buildDict(str string) map[int]string {

	//MAP
	dict := make(map[int]string)

	//Finding defined strings
	dictstring := strings.Split(str, " ")
	for _, str := range dictstring {
		switch len(str) {
		//one
		case 2:
			dict[1] = str
		//seven
		case 3:
			dict[7] = str
		//four
		case 4:
			dict[4] = str
		//eight
		case 7:
			dict[8] = str
		}
	}

	//Finding Positions
	countstring := strings.Split(str, "")
	countdict := make(map[rune]int)

	for _, str := range countstring {
		switch {
		case str == "a":
			countdict['a']++
		case str == "b":
			countdict['b']++
		case str == "c":
			countdict['c']++
		case str == "d":
			countdict['d']++
		case str == "e":
			countdict['e']++
		case str == "f":
			countdict['f']++
		case str == "g":
			countdict['g']++
		}
	}

	/*
		 000
		1   2
		 333
		4   5
		 666
	*/

	posdict := make(map[int]string)
	for k, v := range countdict {
		if v == 6 {
			posdict[1] = string(k)
		}
		if v == 4 {
			posdict[4] = string(k)
		}
		if v == 9 {
			posdict[5] = string(k)
		}
	}

	posdict[2] = strings.ReplaceAll(dict[1], string(posdict[5]), "")
	posdict[0] = strings.ReplaceAll(dict[7], string(posdict[2]), "")
	posdict[0] = strings.ReplaceAll(posdict[0], string(posdict[5]), "")
	posdict[3] = strings.ReplaceAll(dict[4], string(posdict[1]), "")
	posdict[3] = strings.ReplaceAll(posdict[3], string(posdict[2]), "")
	posdict[3] = strings.ReplaceAll(posdict[3], string(posdict[5]), "")

	rune6 := dict[8]
	for _, v := range posdict {
		rune6 = strings.ReplaceAll(rune6, string(v), "")
	}
	posdict[6] = rune6

	//filling the rest of dict
	/*
		 000
		1   2
		 333
		4   5
		 666
	*/
	dict[0] = string(posdict[0]) + string(posdict[1]) + string(posdict[2]) + string(posdict[4]) + string(posdict[5]) + string(posdict[6])
	dict[2] = string(posdict[0]) + string(posdict[2]) + string(posdict[3]) + string(posdict[4]) + string(posdict[6])
	dict[3] = string(posdict[0]) + string(posdict[2]) + string(posdict[3]) + string(posdict[5]) + string(posdict[6])
	dict[5] = string(posdict[0]) + string(posdict[1]) + string(posdict[3]) + string(posdict[5]) + string(posdict[6])
	dict[6] = string(posdict[0]) + string(posdict[1]) + string(posdict[3]) + string(posdict[4]) + string(posdict[5]) + string(posdict[6])
	dict[9] = string(posdict[0]) + string(posdict[1]) + string(posdict[2]) + string(posdict[3]) + string(posdict[5]) + string(posdict[6])

	//Printing dictonary
	//fmt.Println(dict)
	//fmt.Println(posdict)

	return dict
}
