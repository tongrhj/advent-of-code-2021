package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	index, count := 0, 0
	aStart, bStart, cStart, dStart := 0,0,0,0
	// f, err := os.Open("test.txt")
	f, err := os.Open("input.txt")
	if err != nil {
			fmt.Print("There has been an error!: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		if index % 4 == 0 {
			aStart = num

			if index >= 3 && aStart > bStart {
				// fmt.Println("aStart", aStart, " bStart", bStart)
				count ++
			}
		}
		if (index - 1) % 4 == 0 {
			bStart = num

			if index >= 3 && bStart > cStart {
				// fmt.Println("bStart", bStart, " cStart", cStart)
				count ++
			}
		}
		if (index - 2) % 4 == 0 {
			cStart = num

			if index >= 3 && cStart > dStart {
				// fmt.Println("cStart", cStart, " dStart", dStart)
				count ++
			}
		}
		if (index - 3) % 4 == 0 {
			dStart = num

			if index >= 3 && dStart > aStart {
				// fmt.Println("dStart", dStart, " aStart", aStart)
				count ++
			}
		}

		index++
	}
	fmt.Println("count =", count)
}
