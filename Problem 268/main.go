package main

// Using formula 

// func missingNumber(nums []int) int {
//     n := len(nums)
//     expectedSum := n * (n + 1) / 2
//     actualSum := 0
//     for _, num := range nums {
//         actualSum += num
//     }
//     return expectedSum - actualSum
// }

// 🧑‍💻 Go Code with Step-by-Step XOR Traversal

import (
	"fmt"
)

func missingNumber(nums []int) int {
	n := len(nums)
	xorAll := 0

	// XOR all numbers from 0 to n
	for i := 0; i <= n; i++ {
		fmt.Printf("XORing xorAll(%d) ^ i(%d) = %d\n", xorAll, i, xorAll^i)
		xorAll ^= i
	}

	// XOR all elements in the array
	for _, num := range nums {
		fmt.Printf("XORing xorAll(%d) ^ num(%d) = %d\n", xorAll, num, xorAll^num)
		xorAll ^= num
	}

	return xorAll
}

func main() {
	nums := []int{3, 0, 1}
	fmt.Println("Missing number is:", missingNumber(nums))
}



// ✅ Go Code for XOR [1..n]:

// import (
// 	"fmt"
// )

// func missingNumber(nums []int) int {
// 	n := len(nums) + 1 // because one is missing
// 	xorFull := 0
// 	xorNums := 0

// 	for i := 1; i <= n; i++ {
// 		xorFull ^= i
// 	}

// 	for _, num := range nums {
// 		xorNums ^= num
// 	}

// 	return xorFull ^ xorNums
// }

// func main() {
// 	nums := []int{1, 2, 4, 5}
// 	fmt.Println("Missing number is:", missingNumber(nums)) // Should print 3
// }