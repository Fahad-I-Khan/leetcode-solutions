package main

func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    minProd, maxProd, result := nums[0], nums[0], nums[0]

    for i := 1; i < len(nums); i++ {
        if nums[i] < 0 {
            minProd, maxProd = maxProd, minProd
        }

        maxProd = max(nums[i], maxProd*nums[i])
        minProd = min(nums[i], minProd*nums[i])

        result = max(result, maxProd)
    }
    return result 
}

