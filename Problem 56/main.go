package main  // Merge Intervals

import (
    "fmt"
    "sort"
)

func merge(intervals [][]int) [][]int {
    // If there are no intervals, return an empty list
    if len(intervals) == 0 {
        return [][]int{}
    }
	
    // Sort intervals by their start values
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    var merged [][]int
    currentInterval := intervals[0]

    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] <= currentInterval[1] {
            // There is overlap, merge the intervals
            currentInterval[1] = max(currentInterval[1], intervals[i][1])
        } else {
            // No overlap, add the previous interval and update currentInterval
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
    intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
    result := merge(intervals)
    fmt.Println("Merged Intervals:", result) // Outputs: Merged Intervals: [[1 6] [8 10] [15 18]]
}
