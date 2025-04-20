package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n // already valid
	}

	i := 2 // start from index 2

	for j := 2; j < n; j++ {
		if nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}

func main() { // Sample input
	nums := []int{1, 1, 1, 2, 2, 3}

	// Call function
	newLength := removeDuplicates(nums)

	// Print results
	fmt.Printf("New length: %d\n", newLength)
	fmt.Printf("How the array look like: %d\n", nums)
	fmt.Printf("Updated array: %v\n", nums[:newLength])

}
