package main

import (
    "fmt"
    "sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
    // Step 1: Add the new interval to the list
    intervals = append(intervals, newInterval)
    
    // Step 2: Sort the intervals by their start times
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    // Step 3: Merge overlapping intervals
    var merged [][]int
    currentInterval := intervals[0]
    
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] <= currentInterval[1] {
            // There is an overlap, merge the intervals
            currentInterval[1] = max(currentInterval[1], intervals[i][1])
        } else {
            // No overlap, add the previous interval to the result
            merged = append(merged, currentInterval)
            currentInterval = intervals[i]
        }
    }
    // Add the last interval
    merged = append(merged, currentInterval)
    
    return merged
}

// Utility function to find the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    // Example 1
    intervals1 := [][]int{{1, 3}, {6, 9}}
    newInterval1 := []int{2, 5}
    result1 := insert(intervals1, newInterval1)
    fmt.Println("Merged Intervals:", result1) // Outputs: [[1 5] [6 9]]

    // Example 2
    intervals2 := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
    newInterval2 := []int{4, 8}
    result2 := insert(intervals2, newInterval2)
    fmt.Println("Merged Intervals:", result2) // Outputs: [[1 2] [3 10] [12 16]]
}

