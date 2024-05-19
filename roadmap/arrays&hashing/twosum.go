package arraysandhashing

import "sort"

// O(n^2) solution, not pretty but does the job!

func twoSumOn2(nums []int, target int) []int {
	sort.Ints(nums)
	var wanted int
	for i := range nums {
		wanted = target - nums[i]
		for j, num := range nums[i+1:] {
			if num == wanted {
				return []int{i, i + j + 1}
			}
		}
	}
	return nil
}

func twoSumOn(nums []int, target int) []int {
	// Create a map to store the values and their index
	m := make(map[int]int)
	for i, num := range nums {
		// Check if the difference between the target and the current number is in the map
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		// Store the number and its index in the map
		m[num] = i
	}
	return nil
}

