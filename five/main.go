package main

import (
	"cmp"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println("Hello Andrea!")

	readFile, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileString := string(readFile)

	almanac := parseAlmanacMap(fileString)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	channels := make(chan int, len(almanac.initialSeeds))

	go func(wg *sync.WaitGroup) {
		for i := 0; i < len(almanac.initialSeeds); i++ {
			channels <- findSeedLoc(&almanac, almanac.initialSeeds[i])
		}
		wg.Done()
	}(wg)

	lowest := []int{}
	go func(wg *sync.WaitGroup) {
		for i := 0; i < len(almanac.initialSeeds); i++ {
			lowest = append(lowest, <-channels)
		}
		wg.Done()

	}(wg)
	wg.Wait()

	slices.Sort(lowest)
	fmt.Println(lowest[0])

	seedRanges := parseInitialSeedRanges(strings.Split(fileString, "\n\n")[0][7:])
	part2 := reverseFindSeedLoc(&almanac, seedRanges)
	fmt.Println(part2)

}

func parseAlmanacMap(almanacRaw string) Almanac {

	fileStringSplit := strings.Split(almanacRaw, "\n\n")

	almanac := Almanac{}

	initialSeedsRaw := fileStringSplit[0][7:]
	initialSeeds := parseInitialSeeds[*Seed](initialSeedsRaw, reflect.TypeOf(Seed{}))
	almanac.initialSeeds = (initialSeeds)

	almanac.seedToSoilMaps = parseCategoryMaps[*Seed, *Soil](fileStringSplit[1], reflect.TypeOf(Seed{}), reflect.TypeOf(Soil{}))
	almanac.soilToFertilizerMaps = parseCategoryMaps[*Soil, *Fertilizer](fileStringSplit[2], reflect.TypeOf(Soil{}), reflect.TypeOf(Fertilizer{}))
	almanac.fertilizerToWaterMaps = parseCategoryMaps[*Fertilizer, *Water](fileStringSplit[3], reflect.TypeOf(Fertilizer{}), reflect.TypeOf(Water{}))
	almanac.waterToLightMaps = parseCategoryMaps[*Water, *Light](fileStringSplit[4], reflect.TypeOf(Water{}), reflect.TypeOf(Light{}))
	almanac.lightToTempMaps = parseCategoryMaps[*Light, *Temp](fileStringSplit[5], reflect.TypeOf(Light{}), reflect.TypeOf(Temp{}))
	almanac.tempToHumidityMaps = parseCategoryMaps[*Temp, *Humidity](fileStringSplit[6], reflect.TypeOf(Temp{}), reflect.TypeOf(Humidity{}))
	almanac.humidityToLocationMaps = parseCategoryMaps[*Humidity, *Location](fileStringSplit[7], reflect.TypeOf(Humidity{}), reflect.TypeOf(Location{}))

	return almanac

}

func parseCategoryMaps[T, S Category](categoryMapsRaw string, sourceCategoryType reflect.Type, destinationCategoryType reflect.Type) AlmanacMaps[T, S] {

	categoryMaps := strings.Split(categoryMapsRaw, "\n")
	aMs := AlmanacMaps[T, S]{}
	for i := 1; i < len(categoryMaps); i++ {
		if string(categoryMaps[i]) == "" {
			continue
		}
		itemsRaw := strings.Split((categoryMaps[i]), " ")

		source := createCategory(reflect.Type(sourceCategoryType))
		sourceNumber, _ := strconv.Atoi(itemsRaw[1])
		(source).setNumber(sourceNumber)

		destination := createCategory(reflect.Type(destinationCategoryType))
		destinationNumber, _ := strconv.Atoi(itemsRaw[0])
		(destination).setNumber(destinationNumber)

		rangeLen, _ := strconv.Atoi(itemsRaw[2])

		am := AlmanacMapItem[T, S]{source.(T), destination.(S), rangeLen}

		aMs.maps = append(aMs.maps, am)
	}

	return aMs
}

func createCategory(categoryType reflect.Type) Category {
	return reflect.New(categoryType).Interface().(Category)
}

func parseInitialSeedRanges(categoryNumbersRaw string) []SeedRange {
	seedRangesRawList := strings.Split(categoryNumbersRaw, " ")
	var seedRanges []SeedRange

	for i, v := range seedRangesRawList {
		value, _ := strconv.Atoi(v)

		if i%2 == 0 {
			var seedRange SeedRange
			seedRange.start = value
			seedRanges = append(seedRanges, seedRange)
		} else if i%2 == 1 {
			lastNumber := &(seedRanges[len(seedRanges)-1])
			lastNumber.rangeLen = value
		}
	}
	return seedRanges

}

func parseInitialSeeds[T Category](categoryNumbersRaw string, categoryType reflect.Type) []T {
	var categoryNumbers []T

	categoryNumbersRawList := strings.Split(categoryNumbersRaw, " ")
	for _, v := range categoryNumbersRawList {
		value, _ := strconv.Atoi(v)

		categoryNumber := createCategory(reflect.Type(categoryType))
		(categoryNumber).setNumber(value)

		categoryNumbers = append(categoryNumbers, categoryNumber.(T))
	}

	return categoryNumbers
}

type Almanac struct {
	initialSeeds           []*Seed
	seedToSoilMaps         AlmanacMaps[*Seed, *Soil]
	soilToFertilizerMaps   AlmanacMaps[*Soil, *Fertilizer]
	fertilizerToWaterMaps  AlmanacMaps[*Fertilizer, *Water]
	waterToLightMaps       AlmanacMaps[*Water, *Light]
	lightToTempMaps        AlmanacMaps[*Light, *Temp]
	tempToHumidityMaps     AlmanacMaps[*Temp, *Humidity]
	humidityToLocationMaps AlmanacMaps[*Humidity, *Location]
}

type AlmanacMapItem[T, S Category] struct {
	sourceStart      T
	destinationStart S
	rangeLen         int
}

type AlmanacMaps[T, S Category] struct {
	maps []AlmanacMapItem[T, S]
}

type Category interface {
	getNumber() int
	setNumber(int)
}

type Seed struct{ number int }
type Soil struct{ number int }
type Fertilizer struct{ number int }
type Water struct{ number int }
type Light struct{ number int }
type Temp struct{ number int }
type Humidity struct{ number int }
type Location struct{ number int }

type SeedRange struct {
	start    int
	rangeLen int
}

func (s *Seed) setNumber(value int) {
	s.number = value
}

func (s *Soil) setNumber(value int) {
	s.number = value
}

func (s *Fertilizer) setNumber(value int) {
	s.number = value
}

func (s *Water) setNumber(value int) {
	s.number = value
}

func (s *Light) setNumber(value int) {
	s.number = value
}

func (s *Temp) setNumber(value int) {
	s.number = value
}

func (s *Humidity) setNumber(value int) {
	s.number = value
}

func (s *Location) setNumber(value int) {
	s.number = value
}

func (s Seed) getNumber() int {
	return s.number
}

func (s Soil) getNumber() int {
	return s.number
}

func (s Fertilizer) getNumber() int {
	return s.number
}

func (s Water) getNumber() int {
	return s.number
}

func (s Light) getNumber() int {
	return s.number
}

func (s Temp) getNumber() int {
	return s.number
}

func (s Humidity) getNumber() int {
	return s.number
}

func (s Location) getNumber() int {
	return s.number
}

func (A *AlmanacMaps[T, S]) findNext(source Category, destinationType reflect.Type) S {
	destinationValue := createCategory(reflect.Type(destinationType))

	for _, mapItem := range (*A).maps {
		minNumber := mapItem.sourceStart.getNumber()
		maxNumber := mapItem.sourceStart.getNumber() + mapItem.rangeLen
		destinationNumberForSet := mapItem.destinationStart.getNumber()

		if (source.getNumber() >= minNumber) && (source.getNumber() < maxNumber) {
			difference := source.getNumber() - minNumber
			destinationValue.setNumber(destinationNumberForSet + difference)
			return destinationValue.(S)
		}

	}

	(destinationValue).setNumber(source.getNumber())

	return destinationValue.(S)
}

func (A *AlmanacMaps[T, S]) reverseFindNext(source S, destinationType reflect.Type) T {
	destinationValue := createCategory(reflect.Type(destinationType))

	for _, mapItem := range (*A).maps {
		minNumber := mapItem.destinationStart.getNumber()
		maxNumber := mapItem.destinationStart.getNumber() + mapItem.rangeLen
		destinationNumberForSet := mapItem.sourceStart.getNumber()

		if (source.getNumber() >= minNumber) && (source.getNumber() < maxNumber) {
			difference := source.getNumber() - minNumber
			destinationValue.setNumber(destinationNumberForSet + difference)
			return destinationValue.(T)
		}

	}

	(destinationValue).setNumber(source.getNumber())

	return destinationValue.(T)
}

func findSeedLoc(almanac *Almanac, seed *Seed) int {
	soil := (*almanac).seedToSoilMaps.findNext(seed, reflect.TypeOf(Soil{}))
	fertilizer := (*almanac).soilToFertilizerMaps.findNext(soil, reflect.TypeOf(Fertilizer{}))
	water := (*almanac).fertilizerToWaterMaps.findNext(fertilizer, reflect.TypeOf(Water{}))
	light := (*almanac).waterToLightMaps.findNext(water, reflect.TypeOf(Light{}))
	temp := (*almanac).lightToTempMaps.findNext(light, reflect.TypeOf(Temp{}))
	humidity := (*almanac).tempToHumidityMaps.findNext(temp, reflect.TypeOf(Humidity{}))
	location := (*almanac).humidityToLocationMaps.findNext(humidity, reflect.TypeOf(Location{}))

	return (*location).getNumber()
}

func reverseFindSeedLoc(almanac *Almanac, seeds []SeedRange) int {

	slices.SortFunc(seeds, func(a, b SeedRange) int {
		return cmp.Compare(a.start, b.start)
	})

	index := 0
	for {

		location := Location{index}
		var currentSeedVal int
		currentSeedVal = findSeed(almanac, &location)

		if slices.ContainsFunc(seeds, func(a SeedRange) bool {
			return currentSeedVal >= a.start && currentSeedVal < (a.start+a.rangeLen)
		}) {
			return index
		}

		index += 1
	}
}

func findSeed(almanac *Almanac, location *Location) int {
	humidity := almanac.humidityToLocationMaps.reverseFindNext((location), reflect.TypeOf(Humidity{}))
	temp := almanac.tempToHumidityMaps.reverseFindNext(humidity, reflect.TypeOf(Temp{}))
	light := almanac.lightToTempMaps.reverseFindNext(temp, reflect.TypeOf(Light{}))
	water := almanac.waterToLightMaps.reverseFindNext(light, reflect.TypeOf(Water{}))
	fertilizer := almanac.fertilizerToWaterMaps.reverseFindNext(water, reflect.TypeOf(Fertilizer{}))
	soil := almanac.soilToFertilizerMaps.reverseFindNext(fertilizer, reflect.TypeOf(Soil{}))
	seed := almanac.seedToSoilMaps.reverseFindNext(soil, reflect.TypeOf(Seed{}))

	return seed.number
}
