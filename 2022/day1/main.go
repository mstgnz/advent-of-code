package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var list []int

// https://adventofcode.com/2022/day/1
func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reader := bufio.NewScanner(file)

	calories := partOne(reader)
	fmt.Printf("Part One Answer: %d\n", calories)

	calories = partTwo(reader)
	fmt.Printf("Part Two Answer: %d\n", calories)
}

func partTwo(reader *bufio.Scanner) int {
	calories := 0
	sort.Ints(list)
	calories += list[len(list)-1]
	calories += list[len(list)-2]
	calories += list[len(list)-3]
	return calories
}

func partOne(reader *bufio.Scanner) int {
	calories := 0
	sum := 0
	for reader.Scan() {
		line := reader.Text()
		if len(line) == 0 {
			list = append(list, sum)
			if calories < sum {
				calories = sum
			}
			sum = 0
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				continue
			}
			sum += num
		}
	}
	return calories
}
