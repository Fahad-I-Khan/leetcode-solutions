package main

// this one i soleved 
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

//This one from leetcode solutions 

// func containsDuplicate(nums []int) bool {
//     m := make(map[int]bool, len(nums))
    
//     for i := range nums {
//         if m[nums[i]] {
//             return true
//         }
//         m[nums[i]] = true
//     }
//     return false 
// }

// This solution is from leetcode but without using hashMap

// func containsDuplicate(nums []int) bool {
//     l := len(nums)
//     i := 0
//     sorted := true
//     for sorted {
//         for i < l-1 {
//             if nums[i] <= nums[i+1] {
//                 if nums[i] == nums[i+1] {
//                     return true
//                 }
//                 i++
//             } else {
//                 nums[i],nums[i+1] = nums[i+1],nums[i]
//                 sorted = false
//             }
//         }
//         if !sorted {
//             i = 0
//             sorted = true
//         } else {
//             break
//         } 
//     }
//     return false
// }