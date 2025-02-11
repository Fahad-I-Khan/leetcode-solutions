package main

import (
    "fmt"
    "sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }

    // Step 1: Sort intervals by end time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })

    // Step 2: Use a greedy approach to count non-overlapping intervals
    count := 1
    end := intervals[0][1]

    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] >= end {
            count++
            end = intervals[i][1]
        }
    }

    // The number of intervals to remove is the total number minus the number of non-overlapping intervals
    return len(intervals) - count
}

func main() {
    // Example
    intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
    result := eraseOverlapIntervals(intervals)
    fmt.Println("Minimum number of intervals to remove:", result) // Outputs: 1
}
