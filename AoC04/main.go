package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	task01()
	//task02()

}

func task01() {
	file, err := os.Open("./input04.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var isFirstLine bool = true
	var bingoNumList []int
	var lineCount int
	var boardCount int
	var board = [5][5]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	boardList := make(map[int][5][5]int)
	for scanner.Scan() {
		line := scanner.Text()
		//Reading first Line with Bingo numbers.
		if isFirstLine {
			strs := strings.Split(scanner.Text(), ",")
			bingoNumList = make([]int, len(strs))
			for i, str := range strs {
				bingoNumList[i] = convertInt(str)
			}
			isFirstLine = false
		} else {
			if len(line) > 0 {
				//Trimming double and leading whitespaces
				tmpline := strings.ReplaceAll(line, "  ", " ")
				tmpline = strings.TrimSpace(tmpline)
				//Splitting Board String
				slc := strings.Split(tmpline, " ")
				for i := range slc {
					board[lineCount][i] = convertInt(slc[i])
				}
				lineCount++
			} else {
				if lineCount > 0 {
					boardList[boardCount] = board
					board = [5][5]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
					fmt.Println("BingoBoard end")
					lineCount = 0
					boardCount++
				}
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	var drawnNumbers []int
	isBingo := false
	var winBoard [5][5]int
	var winNum int

	for i := 0; !isBingo; i++ {

		drawnNumbers = append(drawnNumbers, bingoNumList[i])

		for _, board := range boardList {

			if checkVert(board, drawnNumbers) || checkHori(board, drawnNumbers) {
				winNum = bingoNumList[i]
				winBoard = board
				isBingo = true
			}
		}

	}

	calcWin(winNum, winBoard, drawnNumbers)

}

func calcWin(winNum int, board [5][5]int, drawnNumbers []int) {
	var sum int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !isIntInArray(board[i][j], drawnNumbers) {
				sum += board[i][j]
			}
		}
	}

	fmt.Println(sum * winNum)

}

func task02() {
	file, err := os.Open("./input04.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func convertInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}

func isIntInArray(x int, array []int) bool {

	for _, num := range array {
		if x == num {
			return true
		}
	}
	return false
}

func checkHori(board [5][5]int, drawnNumbers []int) bool {
	for i := 0; i < 5; i++ {
		if isIntInArray(board[i][0], drawnNumbers) &&
			isIntInArray(board[i][1], drawnNumbers) &&
			isIntInArray(board[i][2], drawnNumbers) &&
			isIntInArray(board[i][3], drawnNumbers) &&
			isIntInArray(board[i][4], drawnNumbers) {
			return true
		}
	}
	return false
}

func checkVert(board [5][5]int, drawnNumbers []int) bool {
	for i := 0; i < 5; i++ {
		if isIntInArray(board[0][i], drawnNumbers) &&
			isIntInArray(board[1][i], drawnNumbers) &&
			isIntInArray(board[2][i], drawnNumbers) &&
			isIntInArray(board[3][i], drawnNumbers) &&
			isIntInArray(board[4][i], drawnNumbers) {
			return true
		}
	}
	return false
}
