package answer

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func (r Range) idInRange(id int) bool {
	return id >= r.start && id <= r.end
}

func MergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return ranges
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	merged := []Range{ranges[0]}

	for _, curr := range ranges[1:] {
		last := &merged[len(merged)-1]
		if curr.start <= last.end {
			if curr.end > last.end {
				last.end = curr.end
			}
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}

func Parse(filePath string) ([]Range, []int) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	ranges := []Range{}
	ingredients := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			startEnd := strings.SplitN(line, "-", 2)
			start, _ := strconv.Atoi(startEnd[0])
			end, _ := strconv.Atoi(startEnd[1])
			ranges = append(ranges, Range{start, end})
		} else {
			id, _ := strconv.Atoi(line)
			ingredients = append(ingredients, id)
		}
	}

	return MergeRanges(ranges), ingredients
}

func Compute(ranges []Range, ingredients []int) int {
	numFresh := 0

	for _, id := range ingredients {
		for _, r := range ranges {
			if r.idInRange(id) {
				numFresh++
			}
		}
	}

	return numFresh
}
