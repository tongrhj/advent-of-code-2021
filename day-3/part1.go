package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	index := 0
	var positions []int
	// f, err := os.Open("test.txt")
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Print("There has been an error!: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineStr := scanner.Text()
		// fmt.Println("lineStr", lineStr)
		lineLen := len(lineStr)
		if index == 0 {
			positions = make([]int, lineLen)
		}
		for i := 0; i < lineLen; i++ {
			bit, _ := strconv.Atoi(lineStr[i:i+1])
			// fmt.Println("bit", bit, i)
			positions[i] += bit
		}
		index++
	}
	fileLen := index + 1
	mostCommonThreshold := (fileLen - 1)/2

	var gammaSb strings.Builder
	var epsilonSb strings.Builder
	for i := 0; i < len(positions); i++ {
		// fmt.Println("Coount of index", i, " is ", positions[i])
		if positions[i] > mostCommonThreshold {
			gammaSb.WriteString("1")
			epsilonSb.WriteString("0")
		} else {
			gammaSb.WriteString("0")
			epsilonSb.WriteString("1")
		}
	}

	gamma, _ := strconv.ParseInt(gammaSb.String(), 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonSb.String(), 2, 64)

	// fmt.Println(gamma, epsilon)
	fmt.Println("poewr consumption =", gamma * epsilon)
}
