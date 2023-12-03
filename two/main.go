package main

import (
	"bufio"
	"fmt"
	"log"
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
	return
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var part1Input []game

	for fileScanner.Scan() {
		part1Input = append(part1Input, parseGame(fileScanner.Text()))
	}

	fmt.Println(part1(part1Input))

}

func part1(games []game) int {
	maxBlue := 14
	maxRed := 12
	maxGreen := 13
	total := 0

	for index, game := range games {
		if game.maxBlue <= maxBlue && game.maxGreen <= maxGreen && game.maxRed <= maxRed {
			total = total + index + 1
		}
	}

	return total

}

func parseGame(input string) game {
	subsets := strings.SplitAfterN(input, ":", 2)[1]
	game := game{maxRed: 0, maxBlue: 0, maxGreen: 0}

	for _, subset := range strings.Split(subsets, ";") {
		for _, cubes := range strings.Split(subset, ",") {

			cube := strings.Split(cubes[1:], " ")
			numberCubes, err := strconv.Atoi(cube[0])
			if err != nil {
				log.Fatal("Could not convert number of cubes to integer")
			}
			switch cube[1] {
			case "red":
				game.maxRed = slices.Max([]int{game.maxRed, numberCubes})
			case "green":
				game.maxGreen = slices.Max([]int{game.maxGreen, numberCubes})
			case "blue":
				game.maxBlue = slices.Max([]int{game.maxBlue, numberCubes})
			}
		}
	}
	return game
}

type game struct {
	maxRed   int
	maxGreen int
	maxBlue  int
}
