package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func is_number(ch string) bool {
	if ch == "1" || ch == "2" || ch == "3" || ch == "4" || ch == "5" || ch == "6" || ch == "7" || ch == "8" || ch == "9" {
		return true
	}
	return false
}

// PART TWO

func convert_to_number(number string) string {
	conversion_map := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4ur",
		"five":  "f5ve",
		"six":   "s6x",
		"seven": "s7ven",
		"eight": "e8ight",
		"nine":  "n9ne",
	}
	val, ok := conversion_map[number]
	if ok {
		return val
	}
	return number
}

func has_prefix(line string, index int) (bool, int) {
	prefixes := map[string]int{
		"one":  1,
		"two":  2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six":  6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}

	for prefix, number := range prefixes {
		if strings.HasPrefix(line[index:], prefix) {
			return true, number
		}
	}

	return false, 0
}

func get_numbers(line string) []int {
	var numbers []int
	for _, c := range line {
		ch := string(c)
		if is_number(ch) {
			number, _ := strconv.Atoi(ch)
			numbers = append(numbers, number)
		}
	}
	return numbers
}
func get_numbers_2(line string) []int {
	var numbers []int
	for i, c := range line {
		ch := string(c)
		if is_number(ch) {
			number, _ := strconv.Atoi(ch)
			numbers = append(numbers, number)
		} else {
			has_p, number := has_prefix(line, i)
			if has_p {
				numbers = append(numbers, number)
			}
		}

	}
	return numbers
}


func main() {
	readFile, err := os.Open("1.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result_1 := 0
	result_2 := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println("Line: ", line)
		numbers := get_numbers(line)
		result_1 += numbers[0]*10 + numbers[len(numbers)-1]

		numbers_2 := get_numbers_2(line)
		fmt.Println("Line 2: ", line, "\n\n")
		result_2 += numbers_2[0]*10 + numbers_2[len(numbers_2)-1]
	}
	fmt.Println("Result Part 1: ", result_1)
	fmt.Println("Result Part 2: ", result_2)
}
