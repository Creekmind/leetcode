package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	storage := make(map[int]int)
	for i, n := range nums {
		k := target - n
		if v, ok := storage[k]; ok {
			return []int{v, i}
		}
		storage[n] = i
	}
	return nil
}

func main() {
	indicies := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println("Result", indicies)
}
