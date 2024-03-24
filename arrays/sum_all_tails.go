package main

func SumAllTails(s ...[]int) []int {
	var sums []int
	for _, numbers := range s {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
