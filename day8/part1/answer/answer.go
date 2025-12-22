package answer

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Parse(filePath string) []Position {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	positions := []Position{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xyz := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(xyz[0])
		y, _ := strconv.Atoi(xyz[1])
		z, _ := strconv.Atoi(xyz[2])
		positions = append(positions, Position{x, y, z})
	}

	return positions
}

type Position struct {
	x int
	y int
	z int
}

func distance(a, b Position) float64 {
	return math.Sqrt(
		math.Pow(float64(a.x-b.x), 2) +
			math.Pow(float64(a.y-b.y), 2) +
			math.Pow(float64(a.z-b.z), 2),
	)
}

func distanceKey(a, b Position) string {
	return fmt.Sprintf(
		"%d,%d,%d-%d,%d,%d",
		a.x, a.y, a.z, b.x, b.y, b.z)
}

func Compute(positions []Position, numConnections int) int {
	distances := map[float64]string{}
	for i := 0; i < len(positions)-1; i++ {
		a := positions[i]
		for j := i + 1; j < len(positions); j++ {
			b := positions[j]
			distances[distance(a, b)] = distanceKey(a, b)
		}
	}

	sortedDistances := make([]float64, 0, len(distances))
	for k := range distances {
		sortedDistances = append(sortedDistances, k)
	}

	sort.Float64s(sortedDistances)

	circuits := [][]string{}

	for i := range numConnections {
		abString := distances[sortedDistances[i]]
		abArr := strings.Split(abString, "-")
		a := abArr[0]
		b := abArr[1]

		aCircuitIdx := -1
		bCircuitIdx := -1

		for i := range circuits {
			if slices.Contains(circuits[i], a) {
				aCircuitIdx = i
			}
			if slices.Contains(circuits[i], b) {
				bCircuitIdx = i
			}
		}

		if aCircuitIdx != -1 && aCircuitIdx == bCircuitIdx {
			continue
		}

		if aCircuitIdx != -1 && bCircuitIdx != -1 {
			circuits[aCircuitIdx] = append(circuits[aCircuitIdx], circuits[bCircuitIdx]...)
			circuits = append(circuits[:bCircuitIdx], circuits[bCircuitIdx+1:]...)
			continue
		}

		if aCircuitIdx != -1 {
			circuits[aCircuitIdx] = append(circuits[aCircuitIdx], b)
			continue
		}

		if bCircuitIdx != -1 {
			circuits[bCircuitIdx] = append(circuits[bCircuitIdx], a)
			continue
		}

		circuits = append(circuits, []string{a, b})
	}

	slices.SortFunc(circuits, func(a, b []string) int {
		if len(a) < len(b) {
			return 1
		}

		if len(a) > len(b) {
			return -1
		}

		return 0
	})

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}
