package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left+1, right+1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}


func main()  {
	numbers := []int{2,7,11,15}
	target := 9
	result := twoSum(numbers, target)
	fmt.Println("Two Sum II: ", result)

}