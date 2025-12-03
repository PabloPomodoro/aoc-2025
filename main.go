package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	idsSum := 0

	data, _ := os.ReadFile("data.txt")
	scopes := strings.SplitSeq(string(data), ",")

	for scope := range scopes {
		start, _ := strconv.Atoi(strings.Split(scope, "-")[0])
		end, _ := strconv.Atoi(strings.Split(scope, "-")[1])

		for id := start; id <= end; id++ {

			idString := strconv.Itoa(id)
			idStringLength := len(idString)

			if idStringLength < 2 {
				continue
			}

			singleDigit := true
			for i := 1; i < idStringLength; i++ {

				if idString[0] != idString[i] {
					singleDigit = false
					break
				}
			}
			if singleDigit {
				idsSum += id
				continue
			}

			if idStringLength < 4 {
				continue
			}

			repeating := false
			for j := 2; j < idStringLength; j++ {
				segment := idString[:j]

				segmentLength := len(segment)

				if idStringLength%segmentLength != 0 {
					continue
				}

				var nextSegment string
				repeating = true
				for k := segmentLength; k < idStringLength; k += segmentLength {
					nextSegment = idString[k : k+segmentLength]

					if segment != nextSegment {
						repeating = false
						break
					}
				}
				if repeating {
					break
				}
			}
			if repeating {
				idsSum += id
			}
		}
	}
	fmt.Println(idsSum)
}
