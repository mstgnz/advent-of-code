package main

import (
	"bufio"
	"fmt"
	"os"
)

// #1 https://adventofcode.com/2015/day/1
func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reader := bufio.NewScanner(file)

	partOne, partTwo := partOneAndTwo(reader)
	fmt.Println("Part 1:", partOne)
	fmt.Println("Part 2:", partTwo)
}

func partOneAndTwo(reader *bufio.Scanner) (int, int) {
	floor := 0
	sequence := 1
	position := 0
	for reader.Scan() {
		line := reader.Text()
		for i := 0; i < len(line); i++ {
			if line[i] == '(' {
				floor++
			} else {
				floor--
				if floor == -1 && position == 0 {
					position = sequence
				}
			}
			sequence++
		}
	}
	return floor, position
}
