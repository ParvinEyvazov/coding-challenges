// https://leetcode.com/problems/two-sum/

package main

import "fmt"

func main() {

	nums := []int{3, 3}
	target := 6

	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {

	indexes := []int{}

Exit:
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				indexes = append(indexes, i, j)
				continue Exit
			}
		}
	}

	return indexes
}
