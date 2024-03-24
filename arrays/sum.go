package main

func Sum(source []int) int {
	var ans int
	for _, number := range source {
		ans += number
	}
	return ans
}
