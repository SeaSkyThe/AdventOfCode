package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type processed_line struct{
  card_num int
  winning_numbers []int
  own_numbers []int
}

func process_line(line string) processed_line{
  // Getting card numbers
  process_line := strings.Split(line, ":")
  card_number, _ := strconv.Atoi(strings.Split(process_line[0], " ")[1])

  numbers := strings.Split(process_line[1], "|")
  winning_numbers := strings.Split(numbers[0], " ")
  own_numbers := strings.Split(numbers[1], " ")
  
  proc_line := processed_line{card_num: card_number, winning_numbers: []int{}, own_numbers: []int{}}
  for _, n := range winning_numbers {
    if(n == ""){
      continue
    }
    converted_num, _ := strconv.Atoi(string(n))
    proc_line.winning_numbers = append(proc_line.winning_numbers, converted_num)
  }
  for _, n := range own_numbers {
    if(n == "") {
      continue
    }
    converted_num, _ := strconv.Atoi(string(n))
    proc_line.own_numbers = append(proc_line.own_numbers, converted_num)
  }
  
  return proc_line
}

func contains(arr []int, val int) bool{
  for _, n := range arr{
    if(val == n){
      return true
    }
  }
  return false
}

func evaluate_line(pl processed_line) int {
  total := 0
  for _, n := range pl.own_numbers{
    if(contains(pl.winning_numbers, n)){
      if(total == 0){
        total = 1
      }else{ 
        total = total*2
      }
    }
  }
  return total
}
// MAIN
func main() {
	readFile, err := os.Open("4.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
  total := 0
  for fileScanner.Scan(){
    line := fileScanner.Text()

    processed_line := process_line(line)
    total += evaluate_line(processed_line)
  }
  fmt.Println("Part 1: ", total)
}
