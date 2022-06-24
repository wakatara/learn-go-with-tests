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

func SumAllTails(slices [][]int) []int {
	total := []int{}
	for _, slice := range slices {
		if len(slice) == 0 {
			total = append(total, 0)
		} else {
			total = append(total, Sum(slice[1:]))
		}
	}
	return total
}
