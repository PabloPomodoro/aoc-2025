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
			magnitude := len(strconv.Itoa(id))
			if magnitude%2 != 0 {
				continue
			}

			midway := magnitude / 2
			firstHalf := strconv.Itoa(id)[:midway]
			secondHalf := strconv.Itoa(id)[midway:]

			if firstHalf == secondHalf {
				idsSum += id
			}
		}
	}

	fmt.Println(idsSum)
}
