package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))

}

func twoSum(nums []int, target int) []int {
	for i, k := range nums {
		for n, m := range nums[i+1:] {

			if k+m == target {
				return []int{i, n + i + 1}
			}
		}
	}
	return nil
}
