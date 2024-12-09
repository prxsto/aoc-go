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
	reports := getInput(path)

	valid := checkReports(reports)

	return valid
}

func getInput(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var reports [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Fields(line)

		var report []int
		for _, num := range vals {
			val, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("error converting string: ", err)
			}
			report = append(report, val)
		}

		reports = append(reports, report)

	}
	return reports
}

func checkReports(reports [][]int) int {
	fmt.Println("valid reports: ")
	var valid int
	for _, report := range reports {
		if len(report) < 2 {
			continue
		}

		var ascending bool
		ordered := true
		initialDiff := report[1] - report[0]

		switch {
		case initialDiff == 0:
			continue
		case 1 <= initialDiff && initialDiff <= 3:
			ascending = true
		case -3 <= initialDiff && initialDiff <= -1:
			ascending = false
		default:
			continue
		}

		for i := 1; i < len(report)-1; i++ {
			diff := report[i+1] - report[i]

			if diff == 0 {
				ordered = false
				break
			}

			if ascending && (diff < 1 || diff > 3) {
				ordered = false
				break
			}

			if !ascending && (diff > -1 || diff < -3) {
				ordered = false
				break
			}
		}

		if ordered {
			valid++
		}
	}
	return valid
}
