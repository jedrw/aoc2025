package answer

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int64
	end   int64
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
			start, err := strconv.ParseInt(startEnd[0], 10, 64)
			if err != nil {
				panic(err)
			}

			end, err := strconv.ParseInt(startEnd[1], 10, 64)
			if err != nil {
				panic(err)
			}

			ranges = append(ranges, Range{start, end})
		}
	}

	return ranges
}

func isAllRepeated(s, firstChunk string) bool {
	if len(firstChunk) == 0 {
		return false
	}
	chunkLen := len(firstChunk)
	if len(s)%chunkLen != 0 {
		return false
	}
	totalChunks := len(s) / chunkLen
	return strings.Repeat(firstChunk, totalChunks) == s
}

func Compute(ranges []Range) int64 {
	answer := int64(0)
	for _, r := range ranges {
		for id := r.start; id <= r.end; id++ {
			intString := strconv.FormatInt(id, 10)
			lenInt := len(intString)
			maxLenRepeated := lenInt / 2
			found := false
			for chunkLen := 1; chunkLen <= maxLenRepeated; chunkLen++ {
				firstChunk := intString[:chunkLen]
				if isAllRepeated(intString, firstChunk) {
					answer += id
					found = true
					break
				}
			}

			if found {
				continue
			}
		}
	}

	return answer
}
