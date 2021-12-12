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

var count int

func task01() {
	file, err := os.Open("./input12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	caveNodes := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		slc := strings.Split(line, "-")

		caveNodes[slc[0]] = append(caveNodes[slc[0]], slc[1])
		caveNodes[slc[1]] = append(caveNodes[slc[1]], slc[0])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, element := range caveNodes["start"] {
		route := make([]string, 0)
		route = append(route, "start")
		followRoute(caveNodes, element, make(map[string]bool), route)
	}

	fmt.Println("Number of Routes ", count)

}

func followRoute(caveNodes map[string][]string, nodeName string, visited map[string]bool, route []string) {

	//copying the map to prevent reference access
	copyMap := make(map[string]bool)
	for key, value := range visited {
		copyMap[key] = value
	}

	if strings.ToLower(nodeName) == nodeName {
		copyMap[nodeName] = true
	}
	route = append(route, nodeName)
	for _, element := range caveNodes[nodeName] {
		if element != "end" && element != "start" {
			if strings.ToLower(element) == element {
				if _, isVisited := copyMap[element]; isVisited {
					//Route end, nothing to do anymore
					continue
				}
			}
			followRoute(caveNodes, element, copyMap, route)

		} else if element == "end" {
			route = append(route, "end")
			count++
			//fmt.Println(route)
		}
	}

}

func task02() {
}
