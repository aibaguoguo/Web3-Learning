package main

import (
	"fmt"
)

func main() {
	fmt.Println(findIndex([]int{4, 1, 3, 5, 2, 6}, 7))
}

func findIndex(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		t, ok := m[target-num]
		if ok {
			return []int{i, t}
		}
		m[num] = i
	}
	return nil
}
