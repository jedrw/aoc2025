package answer

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func Parse(filePath string) []Range {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	ranges := []Range{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rangesStrings := strings.Split(line, ",")
		for _, r := range rangesStrings {
			startEnd := strings.SplitN(r, "-", 2)
			start, err := strconv.Atoi(startEnd[0])
			if err != nil {
				panic(err)
			}

			end, err := strconv.Atoi(startEnd[1])
			if err != nil {
				panic(err)
			}

			ranges = append(ranges, Range{start, end})
		}
	}

	return ranges
}

func lenInt(i int) int {
	if i == 0 {
		return 1
	}

	count := 0
	for i != 0 {
		i /= 10
		count++
	}

	return count
}

func Compute(ranges []Range) int {
	answer := 0
	for _, r := range ranges {
		for id := r.start; id <= r.end; id++ {
			lenInt := lenInt(id)
			if lenInt%2 == 1 {
				continue
			}

			intString := strconv.Itoa(id)
			half := lenInt / 2
			first := intString[:half]
			last := intString[half:]
			if last == first {
				answer += id
			}
		}
	}

	return answer
}
