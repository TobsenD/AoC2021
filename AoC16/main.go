package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	start := time.Now()

	file, err := os.Open("./input15.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Task 01 Number: %d\n", task01(line))
		fmt.Printf("Task 02 Number: %d\n", task02(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}

func task01(inputText string) int {

	bits := decodeHex(inputText)
	buffer := bytes.NewBuffer(bits)
	versionSum, _, _ := parsePackets(buffer)

	return versionSum
}

func task02(inputText string) int {
	bits := decodeHex(inputText)
	buf := bytes.NewBuffer(bits)
	_, value, _ := parsePackets(buf)
	return value
}

func parsePackets(buf *bytes.Buffer) (versionSum, value int, err error) {
	if buf.Len() < 11 {
		return 0, 0, errors.New("not enough bytes for another packet")
	}

	version, typeId := parseHeader(buf)
	versionSum = version
	if typeId == 4 {
		value = parseLiteralValue(buf)
	} else {
		ver, val := parseOperatorPacketContents(typeId, buf)
		versionSum += ver
		value = val
	}
	return versionSum, value, nil
}

func parseHeader(buf *bytes.Buffer) (version, typeId int) {
	version = binaryStringInBytesToInt(buf.Next(3))
	typeId = binaryStringInBytesToInt(buf.Next(3))
	return
}

func parseOperatorPacketContents(typeId int, buf *bytes.Buffer) (versionSum, value int) {
	subpacketIndicator, _ := buf.ReadByte()
	var values []int
	if subpacketIndicator == byte(48) {
		subpacketLength := binaryStringInBytesToInt(buf.Next(15))
		subpacketBytes := bytes.NewBuffer(buf.Next(subpacketLength))
		for {
			ver, val, err := parsePackets(subpacketBytes)
			if err != nil {
				break
			}
			versionSum += ver
			values = append(values, val)
		}
	} else {
		subpacketLength := binaryStringInBytesToInt(buf.Next(11))
		for i := 0; i < subpacketLength; i++ {
			ver, val, _ := parsePackets(buf)
			versionSum += ver
			values = append(values, val)
		}
	}
	if len(values) == 0 {
		return versionSum, 0
	}
	switch typeId {
	case 0: // sum values
		for _, v := range values {
			value += v
		}
	case 1: // product
		value = 1
		for _, v := range values {
			value *= v
		}
	case 2: // min
		value, _ = minMax(values)
	case 3: // max
		_, value = minMax(values)
	case 5: // >
		if values[0] > values[1] {
			value = 1
		}
	case 6: // <
		if values[0] < values[1] {
			value = 1
		}
	case 7: // ==
		if values[0] == values[1] {
			value = 1
		}
	}
	return versionSum, value
}

func parseLiteralValue(buf *bytes.Buffer) int {
	var num []byte
	for p, _ := buf.ReadByte(); p == byte(49); p, _ = buf.ReadByte() {
		num = append(num, buf.Next(4)...)
	}
	num = append(num, buf.Next(4)...)
	return binaryStringInBytesToInt(num)
}

func binaryStringInBytesToInt(bin []byte) int {
	i, _ := strconv.ParseInt(string(bin), 2, 0)
	return int(i)
}

func decodeHex(inputText string) []byte {
	var buffer bytes.Buffer
	for _, char := range inputText {
		bits, _ := hexToBinary(string(char))
		buffer.WriteString(bits)
	}

	return buffer.Bytes()
}

func hexToBinary(char string) (string, error) {
	value, err := strconv.ParseUint(char, 16, 0)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%04b", value), nil
}

func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
