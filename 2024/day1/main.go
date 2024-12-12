package main

import (
	"fmt"
	"math"
	"sort"
)

/**
 * Calculate the total distance between two arrays
 * @param l []int
 * @param r []int
 * @return int
 */

func calculateTotalDistance(l, r []int) int {
	sort.Ints(l)
	sort.Ints(r)

	td := 0
	for i := 0; i < len(l); i++ {
		td += int(math.Abs(float64(l[i] - r[i])))
	}

	return td
}

func main() {
	l := []int{3, 4, 2, 1, 3, 3}
	r := []int{4, 3, 5, 3, 9, 3}

	td := calculateTotalDistance(l, r)

	fmt.Println("Total distance: ", td) // 11
}