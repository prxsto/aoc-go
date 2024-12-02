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
	path := "input-user.txt"
	left, right := getInput(path)

	if part2 {
		m := make(map[int]int)

		for _, num := range right {
			m[num] = m[num] + 1
		}

		var sims []int
		for _, num := range left {
			if m[num] != 0 {
				sims = append(sims, num*m[num])
			}
		}

		return sumSlice(sims)
	}

	slices.Sort(left)
	slices.Sort(right)

	var diffs []int
	for i := range left {
		diffs = append(diffs, diff(left[i], right[i]))
	}

	return sumSlice(diffs)
}

func getInput(path string) ([]int, []int) {
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
	return left, right
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
