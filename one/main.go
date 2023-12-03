package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello Andrea!")
	readFile, err := os.OpenFile("input.txt", os.O_RDONLY, 0777)
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	totalPart1 := 0
	totalPart2 := 0

	for fileScanner.Scan() {
		totalPart1 += getTwoDigitVal(fileScanner.Text(), 1)
		totalPart2 += getTwoDigitVal(fileScanner.Text(), 2)
	}
	fmt.Println(totalPart1)
	fmt.Println(totalPart2)
}

func getTwoDigitVal(input string, part int) int {
	first, last := getFirstLastDigits(input, part)
	return 10*first + last

}

func getFirstLastDigits(line string, part int) (int, int) {
	numbers := []string{}
	if part == 1 {
		numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	} else {
		numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	}
	idx := []int{}
	vals := []int{}

	for j := 0; j < len(line); j++ {
		for i := 0; i < len(numbers); i++ {

			index := strings.Index(line[j:], numbers[i])
			if index == -1 {
				continue
			}
			idx = append(idx, j+index)
			vals = append(vals, i)

		}

	}

	min, max := findMinMax(idx, vals)
	var maxIdx int
	var minIdx int
	if part == 1 {
		maxIdx = max
		minIdx = min
	} else {
		maxIdx = max%9 + 9
		minIdx = min%9 + 9
	}
	finalMin, _ := strconv.Atoi(numbers[minIdx])
	finalMax, _ := strconv.Atoi(numbers[maxIdx])
	return finalMin, finalMax
}

func findMinMax(idx []int, vals []int) (int, int) {
	min := -1
	max := -1
	idxMax := slices.Max(idx)
	idxMin := slices.Min(idx)
	for i := 0; i < len(idx); i++ {
		if min != -1 && max != -1 {
			break
		}
		if idx[i] == idxMin && min == -1 {
			min = vals[i]
		}
		if idx[i] == idxMax && max == -1 {
			max = vals[i]
		}
	}
	return min, max
}
