


// Optional: If You Really Want to Remove In-Place with a Loop

// func removeDuplicates(nums []int) int {
//     for i := 0; i < len(nums)-1; {
//         if nums[i] == nums[i+1] {
//             nums = append(nums[:i+1], nums[i+2:]...) // remove one duplicate
//         } else {
//             i++ // only advance if no removal
//         }
//     }
//     return len(nums)
// }

// ⚠️ Note: This works but is not LeetCode's preferred method — they expect you to avoid changing the slice length and instead modify in-place.


// ✅ Fixed and Correct Approach for LeetCode 26 (Two Pointers)


// func removeDuplicates(nums []int) int {
//     if len(nums) == 0 {
//         return 0
//     }

//     // Index to place next unique element
//     i := 0

//     for j := 1; j < len(nums); j++ {
//         if nums[j] != nums[i] {
//             i++
//             nums[i] = nums[j] // overwrite duplicate
//         }
//     }

//     return i + 1 // length of unique part
// }



// 🧠 Code with Timeline/Diagram Comments:

// package main

// import (
// 	"fmt"
// )

// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	i := 0
// 	fmt.Printf("Start: nums = %v\n\n", nums)

// 	for j := 1; j < len(nums); j++ {
// 		fmt.Printf("Step %d:\n", j)
// 		fmt.Printf("i = %d, j = %d\n", i, j)
// 		fmt.Printf("Comparing nums[i]=%d and nums[j]=%d\n", nums[i], nums[j])

// 		if nums[j] != nums[i] {
// 			i++
// 			nums[i] = nums[j]
// 			fmt.Printf("-> Unique! Moved nums[j] to nums[i]: nums = %v\n", nums)
// 		} else {
// 			fmt.Printf("-> Duplicate. Skipped.\n")
// 		}

// 		// Timeline visual
// 		fmt.Printf("   ")
// 		for k := 0; k < len(nums); k++ {
// 			if k == i && k == j {
// 				fmt.Printf("[i=j]%d ", nums[k])
// 			} else if k == i {
// 				fmt.Printf("[i]%d ", nums[k])
// 			} else if k == j {
// 				fmt.Printf("[j]%d ", nums[k])
// 			} else {
// 				fmt.Printf(" %d  ", nums[k])
// 			}
// 		}
// 		fmt.Println("\n")
// 	}

// 	fmt.Printf("Finished. New length = %d\n", i+1)
// 	fmt.Printf("Resulting nums = %v\n", nums[:i+1])
// 	return i + 1
// }

// func main() {
// 	nums := []int{1, 1, 2, 2, 3}
// 	removeDuplicates(nums)
// }