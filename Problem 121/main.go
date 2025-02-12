package main

import "fmt"

// maxProfit function returns the maximum profit
func maxProfit(prices []int) int {
    // If the prices array is empty, return 0 profit
    if len(prices) == 0 {
        return 0
    }

    maxProfit := 0
    minPrice := prices[0] // Correct variable name

    for _, price := range prices {
        if price < minPrice {
            minPrice = price // Correct reference to minPrice
        } else {
            profit := price - minPrice

            if profit > maxProfit {
                maxProfit = profit
            }
        }
    }
    return maxProfit
}

func main() {
    prices := []int{7, 1, 5, 3, 6, 4}
    fmt.Println("Max profit:", maxProfit(prices)) // Output should be 5
}

// func maxProfit(prices []int) int {
//     minPrice := 10001
//     profit := 0

//     for _, price := range prices {
//         if price < minPrice {
//             minPrice = price
//         } else {
//             profit = max(profit, price - minPrice)
//         }
//     }

//     return profit
// }