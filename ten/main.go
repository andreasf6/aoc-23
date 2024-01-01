package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Dir int

const (
	Up Dir = iota
	Down
	Left
	Right
)

type Point struct {
	x int
	y int
}

const (
	EMPTY      = '.'
	TOP_LEFT   = 'F'
	BOT_LEFT   = 'L'
	TOP_RIGHT  = '7'
	BOT_RIGHT  = 'J'
	VERTICAL   = '|'
	HORIZONTAL = '-'
)

func main() {
	inputString, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	tileGrid, startingPoint, maxLen := parseInput(string(inputString))
	startNeighbours, startNeighboursDirs := findStartNeighbours(startingPoint, maxLen)

	_, points, steps := getSteps(startNeighbours, startNeighboursDirs, startingPoint, tileGrid, maxLen)

	fmt.Println(steps)
	fmt.Println(getEnclosed(points))

}

func parseInput(inputString string) (map[Point]rune, Point, int) {
	tileGrid, startingPoint := make(map[Point]rune), Point{}
	for y, line := range strings.Fields(string(inputString)) {
		for x, char := range line {
			tileGrid[Point{x, y}] = char

			if char == 'S' {
				startingPoint = Point{x, y}
			}
		}
	}

	return tileGrid, startingPoint, len(strings.Fields(string(inputString))[0])
}

func findStartNeighbours(start Point, maxLen int) ([]Point, []Dir) {
	startNeighbours := []Point{
		{start.x, start.y - 1},
		{start.x, start.y + 1},
		{start.x + 1, start.y},
		{start.x - 1, start.y},
	}
	dirs := []Dir{Down, Up, Left, Right}

	return startNeighbours, dirs
}

func isPointInGrid(maxLen int, point Point) bool {
	if point.x >= 0 && point.x < maxLen && point.y >= 0 && point.y < maxLen {
		return true
	}
	return false
}

func getSteps(startNeighbours []Point, startDirs []Dir, startingPoint Point, tileGrid map[Point]rune, maxLen int) (map[Point]rune, []Point, int) {
	for i, neighbour := range startNeighbours {
		loop, points, found := getLoop(tileGrid, neighbour, startDirs[i], maxLen)
		if found {
			loopLength := len(loop)
			if loopLength%2 == 0 {
				return loop, points, loopLength / 2
			}
			return loop, points, loopLength/2 + 1
		}
	}

	return make(map[Point]rune), []Point{}, 0
}

func getLoop(tileGrid map[Point]rune, point Point, fromDir Dir, maxLen int) (map[Point]rune, []Point, bool) {
	loop := make(map[Point]rune)
	var points []Point
	found := false
	tile := tileGrid[Point{point.x, point.y}]

	if tile == EMPTY {
		return loop, points, false
	}

	loop[point] = tile
	points = append(points, point)

	for {
		point, tile, found, fromDir = findNext(tileGrid, point, maxLen, fromDir)

		if !found {
			return loop, points, false
		}

		loop[point] = tileGrid[point]

		points = append(points, point)

		if tile := tileGrid[Point{point.x, point.y}]; tile == 'S' {
			points = append([]Point{point}, points...)
			loop[point] = 'S'
			return loop, points, true
		}
	}
}

func findNext(tileGrid map[Point]rune, point Point, maxLen int, fromDir Dir) (Point, rune, bool, Dir) {
	if !isPointInGrid(maxLen, point) {
		return point, '.', false, fromDir
	}
	currentTile := tileGrid[point]

	switch currentTile {
	case EMPTY:
		return point, currentTile, false, fromDir
	case 'S':
		return point, currentTile, true, fromDir
	case TOP_LEFT:
		if fromDir == Down {
			return Point{x: point.x + 1, y: point.y}, currentTile, true, Left
		} else if fromDir == Right {
			return Point{x: point.x, y: point.y + 1}, currentTile, true, Up
		}
	case TOP_RIGHT:
		if fromDir == Down {
			return Point{x: point.x - 1, y: point.y}, currentTile, true, Right
		} else if fromDir == Left {
			return Point{x: point.x, y: point.y + 1}, currentTile, true, Up
		}
	case BOT_LEFT:
		if fromDir == Up {
			return Point{x: point.x + 1, y: point.y}, currentTile, true, Left
		} else if fromDir == Right {
			return Point{x: point.x, y: point.y - 1}, currentTile, true, Down
		}
	case BOT_RIGHT:
		if fromDir == Up {
			return Point{x: point.x - 1, y: point.y}, currentTile, true, Right
		} else if fromDir == Left {
			return Point{x: point.x, y: point.y - 1}, currentTile, true, Down
		}
	case VERTICAL:
		if fromDir == Up {
			return Point{x: point.x, y: point.y + 1}, currentTile, true, fromDir
		} else if fromDir == Down {
			return Point{x: point.x, y: point.y - 1}, currentTile, true, fromDir
		}
	case HORIZONTAL:
		if fromDir == Right {
			return Point{x: point.x - 1, y: point.y}, currentTile, true, fromDir
		} else if fromDir == Left {
			return Point{x: point.x + 1, y: point.y}, currentTile, true, fromDir
		}

	}
	return point, currentTile, false, fromDir
}

// Shoelace formula and picks theorem
func getEnclosed(points []Point) float64 {
	area := 0.0
	points = points[0 : len(points)-1]
	n := len(points)

	if n < 3 {
		return 0.0
	}

	for i := 0; i < n; i++ {

		if i == n-1 {
			area += (float64(points[i].x) * float64(points[0].y)) - (float64(points[0].x) * float64(points[i].y))
			continue
		}
		area += ((float64(points[i].x) * float64(points[i+1].y)) - (float64(points[i+1].x) * float64(points[i].y)))

	}
	area = math.Abs(area) / 2.0
	interiorPoints := area + 1 - (float64(n) / 2)

	return interiorPoints
}

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
