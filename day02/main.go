package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func extract_data(line string) (int, []map[string]int){
	first_separation := strings.Split(line, ":")
	game_id := strings.Split(first_separation[0], " ")[1]
	subsets := strings.Split(first_separation[1], ";")

	subsets_count := make([]map[string]int, len(subsets))

	for i := range subsets_count {
		subsets_count[i] = make(map[string]int)
		subsets_count[i]["blue"] = 0
		subsets_count[i]["red"] = 0
		subsets_count[i]["green"] = 0
	}

	for i, subset := range subsets {
		colors_count := strings.Split(subset, ",")
		for _, color_count := range colors_count {
			// GET THE GREEN COUNT
			if strings.Contains(color_count, "green") {
				color_count_int, _ := strconv.Atoi(strings.TrimSpace(strings.Split(color_count, "green")[0]))
				subsets_count[i]["green"] = color_count_int
			}else if strings.Contains(color_count, "blue") {
				color_count_int, _ := strconv.Atoi(strings.TrimSpace(strings.Split(color_count, "blue")[0]))
				subsets_count[i]["blue"] = color_count_int
			}else if strings.Contains(color_count, "red") {
				color_count_int, _ := strconv.Atoi(strings.TrimSpace(strings.Split(color_count, "red")[0]))
				subsets_count[i]["red"] = color_count_int
			}



		}
	}
  int_game_id, _ := strconv.Atoi(game_id)
  return int_game_id, subsets_count
}

func process_game(subsets []map[string]int) bool{
  qnt_blue := 14
  qnt_green := 13
  qnt_red := 12

  for _, subset := range subsets{
    if(subset["green"] > qnt_green){
      return false
    }
    if(subset["red"] > qnt_red){
      return false
    }
    if(subset["blue"] > qnt_blue){
      return false
    }
  }
  return true
}

func main() {
	readFile, err := os.Open("2.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
  total := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
    game_id, subsets := extract_data(line)
    if process_game(subsets){
      total += game_id
    }
	}
  fmt.Println("Part 1: ",total)
}
