package main

import (
	"fmt"
	"math"
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
	races := parseRaces(fileString)
	fmt.Printf("%+v\n", races)
	fmt.Println(countTimesToBeatRecord(&races))

	race := parseRace(fileString)
	fmt.Printf("%+v\n", race)

	fmt.Println(int(countTimesToBeatRecord(&race)))

}

type Race struct {
	time     int
	distance int
}

func parseRaces(stringInput string) []Race {
	racesRaw := strings.Split(stringInput, "\n")
	var races []Race

	for i, raceVal := range strings.Fields(racesRaw[0])[1:] {
		race := Race{}
		time, _ := strconv.Atoi(raceVal)
		distance, _ := strconv.Atoi(strings.Fields(racesRaw[1])[1:][i])
		race.time = time
		race.distance = distance
		races = append(races, race)
	}
	return races
}

func parseRace(stringInput string) []Race {
	racesRaw := strings.Split(stringInput, "\n")
	var races []Race
	time, _ := strconv.Atoi(strings.Join(strings.Fields(racesRaw[0])[1:], ""))
	distance, _ := strconv.Atoi(strings.Join(strings.Fields(racesRaw[1])[1:], ""))

	races = append(races, Race{distance: distance, time: time})
	return races
}

func countTimesToBeatRecord(races *[]Race) float64 {
	totalTimes := 1.0
	for _, race := range *races {
		root1, root2, _ := solveQuadratic(1, float64(-race.time), float64(race.distance))
		totalTime := ((float64(race.time) + 1.0) / 2.0) - math.Ceil(math.Max(root1, root2))
		totalTime = totalTime * float64(-2)
		totalTimes *= totalTime
	}
	return totalTimes
}

func solveQuadratic(a, b, c float64) (float64, float64, bool) {
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return 0, 0, false
	}

	sqrtDiscriminant := math.Sqrt(discriminant)

	root1 := (-b + sqrtDiscriminant) / (2 * a)
	root2 := (-b - sqrtDiscriminant) / (2 * a)

	return root1, root2, true
}
