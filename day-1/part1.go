package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	index, count, lastValue := 0, 0, 0
	f, err := os.Open("input.txt")
	if err != nil {
			fmt.Print("There has been an error!: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		if index != 0 && num > lastValue {
			count++
		}
		lastValue = num
		index++
	}
	fmt.Println("count =", count)
}
