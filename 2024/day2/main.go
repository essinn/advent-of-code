package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * isSafe checks if the given report is safe or not
 * A report is considered safe if the difference between each level is 1 or 3
 * @param report []int
 * @return bool
 */

func isSafe(report []int) bool {
	asc := true
	desc := true

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		if diff < 1 || diff > 3 {
			return false
		}

		if diff < 0 {
			asc = false
		} else if diff > 0 {
			desc = false
		}
	}

	return asc || desc
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var reports [][]int

	fmt.Println("Enter reports (enter twice to break): ")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		levels := []int{}
		for _, n := range strings.Fields(line) {
			level, _ := strconv.Atoi(n)
			levels = append(levels, level)
		}

		reports = append(reports, levels)
	}

	safeCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		}
	}

	fmt.Println("number of safe reports: ", safeCount)
}