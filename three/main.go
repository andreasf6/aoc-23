package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
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

	input := Input{symbols: []Symbol{}, nums: []Num{}, partNums: make(map[int]int)}
	lineIndex := 0
	numId := 0
	for fileScanner.Scan() {
		symbs, nums := parseInput(fileScanner.Text(), lineIndex, numId)
		input.symbols = append(input.symbols, symbs...)
		input.nums = append(input.nums, nums...)
		lineIndex += 1
		numId += len(nums)
	}

	getAdjacentSymbols(&input, lineIndex)
	fmt.Println(getTotalNums(keys(input.partNums), &input.nums))

}

type Input struct {
	symbols  []Symbol
	nums     []Num
	partNums map[int]int
}

type Symbol struct {
	lineIndex      int
	positionInLine int
}

type Num struct {
	lineIndex       int
	positionsInLine []int
	value           int
	id              int
}

func parseInput(inputRaw string, lineIndex int, numId int) (symbs []Symbol, nums []Num) {
	symbs = []Symbol{}
	nums = []Num{}
	inputLine := []rune(inputRaw)

	indexInLine := 0
	for indexInLine < len(inputLine) {
		char := inputLine[indexInLine]

		if char != 46 && !unicode.IsDigit(char) {
			symbs = append(symbs, Symbol{positionInLine: indexInLine, lineIndex: lineIndex})
			indexInLine += 1
		} else if unicode.IsDigit(char) {
			valueString := string(char)
			positionsInLine := []int{indexInLine}

			lookAhead := 1
			// build a digit
			for indexInLine+lookAhead < len(inputLine) && unicode.IsDigit(inputLine[indexInLine+lookAhead]) {
				valueString = valueString + string(inputLine[indexInLine+lookAhead])
				positionsInLine = append(positionsInLine, indexInLine+lookAhead)
				lookAhead += 1
			}
			indexInLine += lookAhead

			value, _ := strconv.Atoi(valueString)
			nums = append(nums, Num{positionsInLine: positionsInLine, lineIndex: lineIndex, value: value, id: numId})
			numId += 1
		} else {
			indexInLine += 1
		}

	}

	return
}

func getAdjacentSymbols(inputPointer *Input, maxLineIndex int) {
	input := *inputPointer

	for _, symbol := range input.symbols {
		lineIndex := symbol.lineIndex

		// check in same line
		for _, numSameLine := range getNumsInLine(lineIndex, input.nums) {
			if slices.Contains(numSameLine.positionsInLine, symbol.positionInLine-1) || slices.Contains(numSameLine.positionsInLine, symbol.positionInLine+1) {
				input.partNums[numSameLine.id] = numSameLine.value
			}
		}

		// check line above
		if lineIndex-1 >= 0 {
			for _, numLineAbove := range getNumsInLine(lineIndex-1, input.nums) {
				if slices.Contains(numLineAbove.positionsInLine, symbol.positionInLine-1) || slices.Contains(numLineAbove.positionsInLine, symbol.positionInLine+1) || slices.Contains(numLineAbove.positionsInLine, symbol.positionInLine) {
					input.partNums[numLineAbove.id] = numLineAbove.value
				}
			}
		}

		// check line below
		if lineIndex+1 < maxLineIndex {
			for _, numLineAbove := range getNumsInLine(lineIndex+1, input.nums) {
				if slices.Contains(numLineAbove.positionsInLine, symbol.positionInLine-1) || slices.Contains(numLineAbove.positionsInLine, symbol.positionInLine+1) || slices.Contains(numLineAbove.positionsInLine, symbol.positionInLine) {
					input.partNums[numLineAbove.id] = numLineAbove.value
				}
			}

		}
	}
}

func getNumsInLine(lineIndex int, nums []Num) (numsInLine []Num) {
	for _, num := range nums {
		if num.lineIndex == lineIndex {
			numsInLine = append(numsInLine, num)
		}
	}
	return
}

func getTotalNums(numIds []int, nums *[]Num) (total int) {
	for _, num := range *nums {
		if slices.Contains(numIds, num.id) {
			total += num.value
		}
	}
	return
}

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
