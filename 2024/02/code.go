package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}

	path := "input-user.txt"
	lists := getInput(path)

	// 1. hashmap for each list showing instances of each integer
	// 2. find x * n for each unique int, comparing left to right
	//     - values missing from right list are to be ignored

	return len(lists)
}

func getInput(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Fields(line)

		val1, err := strconv.Atoi(vals[0])
		if err != nil {
			fmt.Println("error converting string: ", err)
		}

		val2, err := strconv.Atoi(vals[1])
		if err != nil {
			fmt.Println("error converting string: ", err)
		}

		left = append(left, val1)
		right = append(right, val2)

	}
	return [][]int{left, right}
}
