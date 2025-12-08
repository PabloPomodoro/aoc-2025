package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	fresh := 0

	data, _ := os.ReadFile("test_data.txt")
	full := strings.Split(string(data), "\n\n")
	scopes := strings.Split(full[0], "\n")

	smallest := 0
	biggest := 0
	vacancies := 0
	vacanciesList := [][2]int{}

	for _, scope := range scopes {

		fmt.Println()
		fmt.Println(scope)

		start, _ := strconv.Atoi(strings.Split(scope, "-")[0])
		end, _ := strconv.Atoi(strings.Split(scope, "-")[1])

		if smallest == 0 {
			smallest = start
		}
		if biggest == 0 {
			biggest = end
		}

		for index, vacany := range vacanciesList {

			fmt.Println(vacanciesList)

			if start <= vacany[0] && end >= vacany[1] {
				vacanciesList = slices.Delete(vacanciesList, index, index+1)
				fmt.Println(vacanciesList)
				fmt.Println()

				continue
			}

			if !(vacany[1] < start || vacany[0] > end) {
				vacanciesList = slices.Delete(vacanciesList, index, index+1)

				if start > vacany[0] && end < vacany[1] {
					vacanciesList = append(vacanciesList, [2]int{vacany[0], start - 1})
					vacanciesList = append(vacanciesList, [2]int{end + 1, vacany[1]})
				} else if start <= vacany[0] && end < vacany[1] {
					vacanciesList = append(vacanciesList, [2]int{end + 1, vacany[1]})
				} else if start > vacany[0] && end > vacany[1] {
					vacanciesList = append(vacanciesList, [2]int{vacany[0], start - 1})
				}
			}
		}

		if start < smallest {
			if end < smallest {
				vacanciesList = append(vacanciesList, [2]int{end, smallest - 1})
			}
			smallest = start
		}

		if end > biggest {
			if start > biggest {
				vacanciesList = append(vacanciesList, [2]int{biggest + 1, start - 1})
			}
			biggest = end
		}
	}

	for m := range vacanciesList {
		vacancies += vacanciesList[m][1] - vacanciesList[m][0] + 1
	}

	fresh = biggest - smallest - vacancies + 1
	fmt.Println()
	fmt.Println(fresh)
}
