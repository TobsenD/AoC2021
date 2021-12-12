package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	task()

}

var count1 int
var count2 int

func task() {
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
		followRouteTask01(caveNodes, element, make(map[string]bool), route)
		followRouteTask02(caveNodes, element, make(map[string]bool), route, "")
	}

	fmt.Println("Number of Routes ", count1)
	fmt.Println("Number of Routes ", count2)

}

func followRouteTask01(caveNodes map[string][]string, nodeName string, visited map[string]bool, route []string) {

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
			followRouteTask01(caveNodes, element, copyMap, route)

		} else if element == "end" {
			route = append(route, "end")
			count1++
			//fmt.Println(route)
			continue
		}
	}

}

func followRouteTask02(caveNodes map[string][]string, nodeName string, visited map[string]bool, route []string, twiceVisitedCave string) {

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
					if twiceVisitedCave == "" {
						// branch out with this small cave
						followRouteTask02(caveNodes, element, copyMap, route, element)
					}
					continue
				}
			}
			followRouteTask02(caveNodes, element, copyMap, route, twiceVisitedCave)

		} else if element == "end" {
			route = append(route, "end")
			count2++
			//fmt.Println(route)
		}
	}
}
