package main

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	dup := make(map[int]bool)
	for _, num := range nums {
		if _, ok := dup[num]; ok {
			return true
		}

		dup[num] = true
	}
	return false
}

