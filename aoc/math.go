package aoc

import "sort"

func Mean(nums []int) float64 {
	var sum int
	for _, num := range nums {
		sum += num
	}

	return float64(sum) / float64(len(nums))
}

func Median(nums []int) int {
	sort.Ints(nums)

	return nums[len(nums)/2]
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Summation(n int) int {
	return (n * (n + 1)) / 2
}
