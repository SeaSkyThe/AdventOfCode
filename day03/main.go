package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type number_pos struct {
	value int
	pos   int
}

func newNumberPos(number string, pos int) number_pos {
	n := number_pos{pos: pos}
	value, _ := strconv.Atoi(number)
	n.value = value
	return n
}

type symbol_pos struct {
	value string
	pos   int
}

func newSymbolPos(symbol string, pos int) symbol_pos {
	n := symbol_pos{pos: pos, value: symbol}
	return n
}

type processed_line struct {
	symbols []symbol_pos
	numbers []number_pos
}

func is_number(chr string) bool {
	_, err := strconv.Atoi(chr)
	return err == nil
}

func process_line(line string) processed_line {
	line += "."
	curr_number := ""
	curr_number_last_char_pos := -1
	var numbers []number_pos
	var symbols []symbol_pos
	for i, ch := range line {
		chr := string(ch)
		if is_number(chr) {
			curr_number_last_char_pos = i
			curr_number += chr

			if i == len(line) {
				numbers = append(numbers, newNumberPos(curr_number, curr_number_last_char_pos-len(curr_number)+1))
			}
		} else {
			if chr != "." {
				symbols = append(symbols, newSymbolPos(chr, i))
			}
			if curr_number != "" {
				numbers = append(numbers, newNumberPos(curr_number, curr_number_last_char_pos-len(curr_number)+1))
			}
			curr_number = ""
		}
	}
	line_output := processed_line{numbers: numbers, symbols: symbols}
	return line_output
}

func check_symbol_in_pos(line processed_line, pos int) bool {
	for _, symbol := range line.symbols {
		if symbol.pos == pos {
			return true
		}
	}
	return false
}

func get_line_total(line *processed_line, previous_line *processed_line, next_line *processed_line) int {
	total := 0
	for _, num := range line.numbers {
		// The possible positions for a symbol to make a number valuable is, in the below or above lines, on the positions that the number occupy in its line
		// and one foward and one previous
		possible_positions := []int{num.pos - 1, num.pos + len(strconv.Itoa(num.value))}
		for i := 0; i < len(strconv.Itoa(num.value)); i++ {
			possible_positions = append(possible_positions, num.pos+i)
		}
		// fmt.Println(num)
		// fmt.Println(possible_positions)

		for _, possible_pos := range possible_positions {
			if check_symbol_in_pos(*line, possible_pos) {
				total += num.value
				break
			} else if previous_line != nil && check_symbol_in_pos(*previous_line, possible_pos) {
				total += num.value
				break
			} else if next_line != nil && check_symbol_in_pos(*next_line, possible_pos) {
				total += num.value
				break
			}
		}
		// fmt.Println(total, "\n\n")
		continue
	}
	return total
}

func get_total(processed_lines []processed_line) int {
	total := 0
	for index, line := range processed_lines {
		// fmt.Println("Line number ", index, ": ", line)
		line_total := 0
		if index == 0 {
			line_total += get_line_total(&line, nil, &processed_lines[index+1])
		} else if index == len(processed_lines)-1 {
			line_total += get_line_total(&line, &processed_lines[index-1], nil)
		} else {
			line_total += get_line_total(&line, &processed_lines[index-1], &processed_lines[index+1])
		}
		// fmt.Println(line_total, "\n")
		total += line_total
	}
	return total
}

// PART 2
//
//
//
func insertIfNotExists(slice []number_pos, newItem number_pos) []number_pos {
   for _, item := range slice {
       if item == newItem {
           return slice
       }
   }
   return append(slice, newItem)
}
func check_number_in_pos(line *processed_line, pos int) *number_pos {
	for _, number := range line.numbers {
		if number.pos <= pos && number.pos+len(strconv.Itoa(number.value)) > pos {
			return &number
		}
	}
	return nil
}

func calculate_gear_ratio(line *processed_line, previous_line *processed_line, next_line *processed_line, pos int) int {
  //fmt.Println("\n\nLine being processed: ", line)
	possible_positions := []int{pos - 1, pos, pos + 1}
	nums := []number_pos{}
	for _, possible_pos := range possible_positions {
		num := check_number_in_pos(line, possible_pos)
		num_prev := check_number_in_pos(previous_line, possible_pos)
		num_next := check_number_in_pos(next_line, possible_pos)
		//fmt.Println("Nums: ", num, num_prev, num_next)
		if num != nil {
			nums = insertIfNotExists(nums, *num)
		}
		if num_prev != nil {
			nums = insertIfNotExists(nums, *num_prev)
		}
		if num_next != nil {
			nums = insertIfNotExists(nums, *num_next)
		}
    //fmt.Println("Nums durante o loop: ", nums)
	}
	if len(nums) != 2 {
		return 0
	}
	num1 := nums[0]
	num2 := nums[1]
	//fmt.Println("Count: ", len(nums))
	//fmt.Println("Num1 e Num2: ", num1, num2, "\n\n")
	return num1.value * num2.value
}

func calculate_total_gear_ratio(lines []processed_line) int {
	total := 0
	for index, line := range lines {
		for _, symbol := range line.symbols {
			if symbol.value == "*" {
				if index == 0 {
					total += calculate_gear_ratio(&line, nil, &lines[index+1], symbol.pos)
				} else if index == len(lines)-1 {
					total += calculate_gear_ratio(&line, &lines[index-1], nil, symbol.pos)
				} else {
					total += calculate_gear_ratio(&line, &lines[index-1], &lines[index+1], symbol.pos)
				}
			}
		}
	}
	return total
}

// MAIN
func main() {
	readFile, err := os.Open("3.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var processed_lines []processed_line
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(len(line))
		proc_line := process_line(line)
		processed_lines = append(processed_lines, proc_line)
	}

	fmt.Println("Part 1: ", get_total(processed_lines))
	fmt.Println("Part 2: ", calculate_total_gear_ratio(processed_lines), "\n")
}
