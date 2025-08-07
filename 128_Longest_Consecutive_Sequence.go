package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	hashset := make(map[int]bool)
	for _, num := range nums {
		hashset[num] = true
	}

	longestSeq := 0

	for num := range hashset {
		// Check if this is the start of a sequence
		if !hashset[num-1] {
			seq := 1
			current := num + 1
			for hashset[current] {
				seq++
				current++
			}
			if seq > longestSeq {
				longestSeq = seq
			}
		}
	}

	return longestSeq
}

func main() {
	inputNums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	res := longestConsecutive(inputNums)
	fmt.Println(res) // Output: 9
}
