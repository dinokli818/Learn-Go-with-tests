package main

func SumAll(s ...[]int) []int {
	var ans []int
	for _, numbers := range s {
		ans = append(ans, Sum(numbers))
	}
	return ans
}
