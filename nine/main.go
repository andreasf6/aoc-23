package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello Andrea!")

	readFile, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileString := string(readFile)
	valuesHistory := parseInput(fileString)

	total := 0
	for _, valueHistory := range valuesHistory {
		total += calculateNextNumber(valueHistory.historyVals)
	}
	fmt.Println(total)

	total2 := 0
	for _, valueHistory := range valuesHistory {
		reverseSlice(valueHistory.historyVals)
		if len(valueHistory.historyVals) < 2 {
			break
		}

		total2 += calculateNextNumber(valueHistory.historyVals)
	}
	fmt.Println(total2)
}

type ValueHistory struct {
	historyVals []int
}

func parseInput(rawInput string) []ValueHistory {
	fileSplit := strings.Split(rawInput, "\n")

	values := []ValueHistory{}
	for _, valueHistoryLine := range fileSplit {
		valueHistory := ValueHistory{}

		for _, valueHistoryNum := range strings.Split(valueHistoryLine, " ") {
			num, _ := strconv.Atoi(valueHistoryNum)
			valueHistory.historyVals = append(valueHistory.historyVals, num)
		}
		values = append(values, valueHistory)

	}
	return values
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func calculateNextNumber(sequence []int) int {
	if allZero(sequence) {
		return sequence[0]
	}

	tmpSequence := make([]int, len(sequence)-1)
	for i := 0; i < len(sequence)-1; i++ {
		tmpSequence[i] = sequence[i+1] - sequence[i]
	}

	return sequence[len(sequence)-1] + calculateNextNumber(tmpSequence)
}

func allZero(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}
