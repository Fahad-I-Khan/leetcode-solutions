package main //350. Intersection of Two Arrays

import "sort" // nput: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]
func intersect(nums1 []int, nums2 []int) []int {
	// Create a map to count the occurrences of each element in nums1
	countMap := make(map[int]int)

	for _, num := range nums1 {
		countMap[num]++
	}

	// Create a slice to store the result
	result := []int{}

	// Iterate through nums2 and find common elements
	for _, num := range nums2 {
		if count, exists := countMap[num]; exists && count > 0 {
			result = append(result, num)
			countMap[num]--
		}
	}

	return result
}

