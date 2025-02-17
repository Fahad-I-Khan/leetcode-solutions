package main  // Container With Most Water

func maxArea(height []int) int {
	j, k := 0, len(height) - 1
	maxArea := 0
 
	for j < k {
		 width := k - j 
		 currentSum := min(height[j], height[k]) * width
 
		 maxArea = max(maxArea, currentSum)
 
		 if height[j] < height[k] {
			 j++
		 } else {
			 k--
		 }
	}
	return maxArea
 }