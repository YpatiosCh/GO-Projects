package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var num int
	arr := []int{}

	for i := 0; i < 12500; i++ {
		fmt.Fscan(os.Stdin, &num)
		arr = append(arr, num)

		low, up := Range(arr)
		fmt.Println(low, up)
	}
}

// Calculate the median
func Median(arr []int) int {
	n := len(arr)
	sort.Ints(arr)

	if n%2 == 1 {
		return (arr[(n-1)/2])
	} else {
		return (arr[n/2] + arr[n/2-1]) / 2
	}
}

// Range calculates the median and returns the lowerbound (median - 45) and upperbound (median + 45) of the range.
func Range(arr []int) (int, int) {
	median := Median(arr)
	low := median - 45
	up := median + 45
	return low, up
}
