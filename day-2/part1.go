package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	depth, position := 0, 0
	// f, err := os.Open("test.txt")
	f, err := os.Open("input.txt")
	if err != nil {
			fmt.Print("There has been an error!: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineStr := scanner.Text()
		x := strings.Split(lineStr, " ")
		direction, amtStr := x[0], x[1]
		amt, _ := strconv.Atoi(amtStr)
		if direction == "forward" {
			position += amt
		}
		if direction == "up" {
			depth -= amt
		}
		if direction == "down" {
			depth += amt
		}
	}
	fmt.Println("total =", depth * position)
}
