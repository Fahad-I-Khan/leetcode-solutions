package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)

	for _, i := range nums1 {
		for _, j := range nums2 {
			m[i+j]++
		}
	}

	result := 0 

	for _, k := range nums3 {
		for _, l := range nums4 {
			sum := k + l
			complement := -sum
			if count, exists := m[complement]; exists {
				result += count
			}
		}
	}

	return result
}

func main()  {
	nums1 := []int{2,2} 
	nums2 := []int{-2,-1} 
	nums3 := []int{-1,2} 
	nums4 := []int{0,2}

	result := fourSumCount(nums1, nums2, nums3, nums4)
	println("Total count:", result)
}