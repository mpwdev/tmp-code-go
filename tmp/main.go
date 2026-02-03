package main

import (
	"fmt"
)

func main() {
	var nums []int
	fmt.Printf("nums: %#v, type: %T, len: %d, cap: %d\n", nums, nums, len(nums), cap(nums))

	nums = append(nums, 1, 2)
	fmt.Printf("nums: %#v, type: %T, len: %d, cap: %d\n", nums, nums, len(nums), cap(nums))

	nums = append(nums, 3, 4, 5)
	fmt.Printf("nums: %#v, type: %T, len: %d, cap: %d\n", nums, nums, len(nums), cap(nums))

	nums = append(nums, 6, 7, 8, 9, 10)
	fmt.Printf("nums: %#v, type: %T, len: %d, cap: %d\n", nums, nums, len(nums), cap(nums))

	nums = append(nums, 11, 12, 13, 14, 15)
}
