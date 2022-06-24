package main

func Sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func SumAll(slices [][]int) []int {
	total := []int{}
	for _, slice := range slices {
		total = append(total, Sum(slice))
	}
	return total
}
