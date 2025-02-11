package main

// This took 4ms
func twoSum(nums []int, target int) []int {

	numMap := make(map[rune]int)

	for i, num := range nums {
		complement := target - num

		if index, found := numMap[rune(complement)]; found {
			return []int{index, i}
		}
		numMap[rune(num)] = i
	}
	return nil
}

// This took 0ms
// func twoSum(nums []int, target int) []int {
//     m := make(map[int]int)
//     res := make([]int, 2)

//     for i := 0; i < len(nums); i++ {
//         num := nums[i]
//         complement := target - nums[i]

//         _, prs := m[complement]
//         if prs {
//             res[0] = m[complement]
//             res[1] = i
//             return res
//         }

//         m[num] = i
//     }
// }

// This took 1ms
// func twoSum(nums []int, target int) []int {
// 	numsLen := len(nums) 
// 	if numsLen < 2 {
// 		return []int{}
// 	}
// After I remove numLen it took 0ms 
// 	if len(nums) < 2 {
// 		return []int{}
// 	}

// 	numToIndex := make(map[int]int)
// 	for idx, num := range nums {
// 		need := target - num
// 		mIdx, ok := numToIndex[need]
// 		if ok {
// 			return []int{mIdx, idx}
// 		}

// 		numToIndex[num] = idx
// 	}
// }

// func main() {
// 	nums := []int{2, 7, 11, 15}
// 	target := 9
// 	result := twoSum(nums, target)
// 	fmt.Println(result) // Output: [0, 1]
// }

// package main

// import (
//     "fmt"
//     "testing"
// )

// func twoSumSolution1(nums []int, target int) []int {
//     m := make(map[int]int)   // Creating the map
//     res := make([]int, 2)    // Creating the result slice

//     for i := 0; i < len(nums); i++ {
//         num := nums[i]
//         complement := target - num
//         _, prs := m[complement] // Checking if the complement exists in the map
//         if prs {
//             res[0] = m[complement] // Store the index of the complement
//             res[1] = i             // Store the current index
//             return res             // Return the result
//         }
//         m[num] = i               // Add the current number and its index to the map
//     }
//     return nil // Return nil if no solution found
// }

// func twoSumSolution2(nums []int, target int) []int {
//     m := new(map[int]int)     // Create a pointer to a nil map
//     *m = make(map[int]int)    // Initialize the map through dereferencing the pointer

//     res := new([]int)         // Create a pointer to a nil slice
//     *res = make([]int, 2)     // Initialize the slice through dereferencing the pointer

//     for i := 0; i < len(nums); i++ {
//         num := nums[i]
//         complement := target - num
//         _, prs := (*m)[complement] // Checking if the complement exists in the map
//         if prs {
//             (*res)[0] = (*m)[complement] // Store the index of the complement
//             (*res)[1] = i                 // Store the current index
//             return *res                   // Return the result
//         }
//         (*m)[num] = i                   // Add the current number and its index to the map
//     }
//     return nil // Return nil if no solution found
// }

// func twoSumSolution3(nums []int, target int) []int {
//     if len(nums) < 2 {
//         return []int{}
//     }

//     numToIndex := make(map[int]int)  // Create the map to store number to index
//     for idx, num := range nums {
//         need := target - num
//         mIdx, ok := numToIndex[need]  // Check if the complement exists in the map
//         if ok {
//             return []int{mIdx, idx}  // If found, return the indices of the pair
//         }
//         numToIndex[num] = idx  // Otherwise, add the current number and its index to the map
//     }

//     return []int{} // Return an empty slice if no pair is found
// }

// // Benchmark functions
// func BenchmarkTwoSumSolution1(b *testing.B) {
//     nums := []int{2, 7, 11, 15, 1, 8, 4, 3, 5, 9, 6} // You can vary this input
//     target := 9
//     b.ResetTimer() // Reset the timer to exclude setup time
//     for i := 0; i < b.N; i++ {
//         twoSumSolution1(nums, target)
//     }
// }

// func BenchmarkTwoSumSolution2(b *testing.B) {
//     nums := []int{2, 7, 11, 15, 1, 8, 4, 3, 5, 9, 6} // You can vary this input
//     target := 9
//     b.ResetTimer() // Reset the timer to exclude setup time
//     for i := 0; i < b.N; i++ {
//         twoSumSolution2(nums, target)
//     }
// }

// func BenchmarkTwoSumSolution3(b *testing.B) {
//     nums := []int{2, 7, 11, 15, 1, 8, 4, 3, 5, 9, 6} // You can vary this input
//     target := 9
//     b.ResetTimer() // Reset the timer to exclude setup time
//     for i := 0; i < b.N; i++ {
//         twoSumSolution3(nums, target)
//     }
// }

// func main() {
//     fmt.Println("Benchmarking...")
//     result := testing.Benchmark(BenchmarkTwoSumSolution1)
//     fmt.Println("Solution 1:", result)
//     result = testing.Benchmark(BenchmarkTwoSumSolution2)
//     fmt.Println("Solution 2:", result)
//     result = testing.Benchmark(BenchmarkTwoSumSolution3)
//     fmt.Println("Solution 3:", result)
// }
