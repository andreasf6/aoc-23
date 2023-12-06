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

	scratchCards := []ScratchCard{}
	for fileScanner.Scan() {
		rawInput := strings.Fields(strings.Split(fileScanner.Text(), ":")[1])
		winNumsRaw, myNumsRaw := rawInput[0:slices.Index(rawInput, "|")], rawInput[slices.Index(rawInput, "|")+1:]

		sC := ScratchCard{winning: []int{}, myNums: [25]int{}, points: 0}
		winningInCard := 0
		for i, myNumber := range myNumsRaw {
			numberVal, _ := strconv.Atoi(myNumber)
			sC.myNums[i] = numberVal
			if slices.Contains(winNumsRaw, myNumber) {
				winningInCard += 1
				sC.winning = append(sC.winning, numberVal)
			}
		}
		if winningInCard > 0 {
			sC.points = 1 << winningInCard >> 1

		}
		scratchCards = append(scratchCards, sC)
	}

	fmt.Println(calculateTotal(&scratchCards))
	fmt.Println(calculateTotalCardNums(&scratchCards))

}

type ScratchCard struct {
	winning []int
	myNums  [25]int
	points  int
}

func calculateTotal(sCs *[]ScratchCard) (total int) {
	for _, sC := range *sCs {
		total += sC.points
	}
	return
}

func calculateTotalCardNums(sCs *[]ScratchCard) (total int) {
	pile := make(map[int]int)
	maxInt := 0
	for i, sC := range *sCs {
		pile[i] += 1
		if sC.points > 0 {

			for k := 0; k < pile[i]; k++ {
				for j := 1; j <= len(sC.winning); j++ {
					if j+i < len(*sCs) {
						pile[j+i] += 1
					}
				}
			}
		}
	}

	for _, v := range pile {
		if maxInt < v {
			maxInt = v
		}
		total += v
	}
	fmt.Println(maxInt)
	return
}
