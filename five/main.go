package main

import (
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

	almanac := parseAlmanacMap(fileString, false)

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

}

func parseAlmanacMap(almanacRaw string, part2 bool) Almanac {

	fileStringSplit := strings.Split(almanacRaw, "\n\n")

	almanac := Almanac{}

	initialSeedsRaw := fileStringSplit[0][7:]
	initialSeeds := parseInitialSeeds[*Seed](initialSeedsRaw, reflect.TypeOf(Seed{}), part2)
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

func parseInitialSeeds[T Category](categoryNumbersRaw string, categoryType reflect.Type, part2 bool) []T {
	var categoryNumbers []T

	categoryNumbersRawList := strings.Split(categoryNumbersRaw, " ")
	for i, v := range categoryNumbersRawList {
		value, _ := strconv.Atoi(v)

		if (part2) && i%2 == 1 {
			lastNumber := categoryNumbers[len(categoryNumbers)-1]
			for i = 1; i <= value; i++ {
				categoryNumber := createCategory(reflect.Type(categoryType))
				(categoryNumber).setNumber(lastNumber.getNumber() + i)

				categoryNumbers = append(categoryNumbers, categoryNumber.(T))

			}
			continue
		}
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

func (A *AlmanacMaps[T, S]) findNext(source T, destinationType reflect.Type) S {
	destinationValue := createCategory(reflect.Type(destinationType))

	for _, mapItem := range (*A).maps {

		minNumber := mapItem.sourceStart.getNumber()
		maxNumber := mapItem.sourceStart.getNumber() + mapItem.rangeLen
		if (source.getNumber() >= minNumber) && (source.getNumber() < maxNumber) {
			difference := source.getNumber() - minNumber
			destinationValue.setNumber(mapItem.destinationStart.getNumber() + difference)
			return destinationValue.(S)
		}

	}

	(destinationValue).setNumber(source.getNumber())

	return destinationValue.(S)
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
