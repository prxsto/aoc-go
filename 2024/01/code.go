package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
	// solve part 1 here

	file, err := os.Open("input-user.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var l1 []int
	var l2 []int

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

		l1 = append(l1, val1)
		l2 = append(l2, val2)

		slices.Sort(l1)
		slices.Sort(l2)
	}

	var diffs []int
	for i := range l1 {
		diffs = append(diffs, diff(l1[i], l2[i]))
	}

	return sumSlice(diffs)
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func sumSlice(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
