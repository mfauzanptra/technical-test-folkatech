package main

import (
	"fmt"
	"sort"
)

func highestNumber(slice []int) int {
	res := 0
	for _, s := range slice {
		if s > res {
			res = s
		}
	}
	return res
}

func sliceContains(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func MaxProfit(prices []int, i int) int {
	res, buy := 0, 0
	tmpSlice := []int{}

	sortedPrices := make([]int, len(prices))
	copy(sortedPrices, prices)
	sort.Ints(sortedPrices)

	lowest := sortedPrices[:i]

	for _, v := range prices {
		if sliceContains(lowest, v) && buy == 0 {
			buy = v
		} else if sliceContains(lowest, v) {
			res += highestNumber(tmpSlice) - buy
			tmpSlice = []int{}
			buy = v
		} else if v == prices[len(prices)-1] {
			tmpSlice = append(tmpSlice, v)
			res += highestNumber(tmpSlice) - buy
		} else {
			tmpSlice = append(tmpSlice, v)
		}
	}
	return res
}

func main() {
	fmt.Println(MaxProfit([]int{4, 11, 2, 20, 59, 80}, 2))
}
