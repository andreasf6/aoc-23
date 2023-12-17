package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	L string
	R string
}

type Network struct {
	nodes map[string]Node
}

func parseInput(rawInput string) ([]string, Network) {
	fileSplit := strings.Split(rawInput, "\n\n")
	instructions := strings.Split(fileSplit[0], "")

	network := Network{nodes: make(map[string]Node)}
	for _, lineFull := range strings.Split(fileSplit[1], "\n") {
		if len(lineFull) < 1 {
			break
		}
		lineSplit := strings.Split(lineFull, " = ")
		var re = regexp.MustCompile(`(?m)(\w+, \w+)`)
		valuesNodes := strings.Split(re.FindString(lineSplit[1]), ", ")

		network.nodes[lineSplit[0]] = Node{L: valuesNodes[0], R: valuesNodes[1]}
	}
	return instructions, network
}

func followInstructions(instructions []string, network *Network, start string, end string) int {
	current := start
	currentInstruction := 0

	for {
		if strings.HasSuffix(current, end) {
			break
		}
		current = network.getNext(current, instructions[currentInstruction%len(instructions)])
		currentInstruction += 1
	}
	return currentInstruction
}

func (N *Network) getNext(val string, instruction string) string {
	switch instruction {
	case "L":
		return N.nodes[val].L
	case "R":
		return N.nodes[val].R
	}
	return ""
}

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	fmt.Println("Hello Andrea!")

	readFile, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileString := string(readFile)
	instructions, network := parseInput(fileString)

	fmt.Println(followInstructions(instructions, &network, "AAA", "ZZZ"))

	part2Start := "A"
	startPossibilities := []string{}
	networkNodes := keys(network.nodes)
	for i := 0; i < len(networkNodes); i++ {
		if strings.HasSuffix(networkNodes[i], part2Start) {
			startPossibilities = append(startPossibilities, networkNodes[i])
		}
	}

}
