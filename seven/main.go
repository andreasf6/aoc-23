package main

import (
	"cmp"
	"fmt"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	hands []Hand
}

type Hand struct {
	cards []Card
	bid   int
}

type Card struct {
	value string
}

func sort(g Game, cardNameVals map[string]int) []Hand {
	slices.SortFunc(g.hands, func(handA, handB Hand) int {
		comparison := cmp.Compare(calculate(handA, cardNameVals), calculate(handB, cardNameVals))
		if comparison == 0 {
			index := 0
			for comparison == 0 {
				comparison = cmp.Compare(cardNameVals[handA.cards[index].value],
					cardNameVals[handB.cards[index].value])
				index++
			}
		}
		return comparison
	})

	return g.hands
}

func parseInput(rawInput string) Game {
	handsString := strings.Split(rawInput, "\n")
	hands := []Hand{}

	for _, handString := range handsString {
		handSplitString := strings.Split(handString, " ")
		if len(handSplitString) < 2 {
			break
		}
		cards := parseCards(handSplitString[0])
		bid, _ := strconv.Atoi(handSplitString[1])
		hands = append(hands, Hand{cards, bid})
	}
	return Game{hands}
}

func parseCards(cardString string) []Card {
	cards := []Card{}
	for _, card := range strings.Split(cardString, "") {
		cards = append(cards, Card{card})
	}
	return cards
}

func calculate(h Hand, cardNameVals map[string]int) int {
	cardValsMap := map[string]int{}
	part2 := false

	if cardNameVals["J"] == 1 {
		part2 = true
	}

originalCards:
	for _, card := range h.cards {
		for cardName := range cardNameVals {
			if card.value == cardName {
				cardValsMap[card.value]++
				continue originalCards
			}
		}
	}

	counts := getCardCounts(cardValsMap, part2)
	slices.Sort(counts)

	switch {
	case reflect.DeepEqual(counts, []int{5}):
		return 7
	case reflect.DeepEqual(counts, []int{1, 4}):
		return 6
	case reflect.DeepEqual(counts, []int{2, 3}):
		return 5
	case reflect.DeepEqual(counts, []int{1, 1, 3}):
		return 4
	case reflect.DeepEqual(counts, []int{1, 2, 2}):
		return 3
	case reflect.DeepEqual(counts, []int{1, 1, 1, 2}):
		return 2
	case reflect.DeepEqual(counts, []int{1, 1, 1, 1, 1}):
		return 1
	default:
		return 0
	}
}

func getCardCounts(cardValsMap map[string]int, part2 bool) []int {
	counts := []int{}
	jIndex := 0
	count := 0
	for k, v := range cardValsMap {
		counts = append(counts, v)
		if k == "J" {
			jIndex = count
		}
		count += 1
	}

	if part2 && cardValsMap["J"] != 0 && cardValsMap["J"] != 5 {
		counts = remove(counts, jIndex)
		counts[slices.Index(counts, slices.Max(counts))] = counts[slices.Index(counts, slices.Max(counts))] + cardValsMap["J"]
	}
	return counts
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	fmt.Println("Hello Andrea!")

	readFile, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileString := string(readFile)
	hands := parseInput(fileString)

	cardNameVals := map[string]int{
		"A": 13, "K": 12, "Q": 11, "J": 10, "T": 9, "9": 8, "8": 7, "7": 6, "6": 5, "5": 4, "4": 3, "3": 2, "2": 1,
	}

	handsPart1 := sort(hands, cardNameVals)

	cardNameVals2 := map[string]int{
		"A": 13, "K": 12, "Q": 11, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2, "J": 1,
	}

	total := 0
	for i, hand := range handsPart1 {
		total += (i + 1) * hand.bid
	}
	fmt.Println(total)

	handsPart2 := sort(hands, cardNameVals2)
	total = 0
	for i, hand := range handsPart2 {
		total += (i + 1) * hand.bid
	}
	fmt.Println(total)

}
