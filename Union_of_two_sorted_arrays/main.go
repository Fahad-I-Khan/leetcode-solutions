package main

import "fmt"

func unionOfSortedArrays(arr1, arr2 []int) []int {
	i, j := 0, 0
	union := []int{}
	
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			if len(union) == 0 || union[len(union)-1] != arr1[i] {
				union = append(union, arr1[i])
			}
			i++
		} else if arr1[i] > arr2[j] {
			if len(union) == 0 || union[len(union)-1] != arr2[j] {
				union = append(union, arr2[j])
			}
			j++
		} else {
			if len(union) == 0 || union[len(union)-1] != arr1[i] {
				union = append(union, arr1[i])
			}
			i++
			j++
		}
	}

	for i < len(arr1) {
		if len(union) == 0 || union[len(union)-1] != arr1[i] {
			union = append(union, arr1[i])
		}
		i++
	}

	for j < len(arr2) {
		if len(union) == 0 || union[len(union)-1] != arr2[j] {
			union = append(union, arr2[j])
		}
		j++
	}

	return union
}

func unionOfSortedArrays2(arr1, arr2 []int) []int {
	set := make(map[int]struct{})
	for _, num := range arr1 {
		set[num] = struct{}{}
	}
	for _, num := range arr2 {
		set[num] = struct{}{}
	}

	union := []int{}
	for num := range set {
		union = append(union, num)
	}

	// Optional: sort the union slice if needed
	// sort.Ints(union)

	return union
}

func main() {
	arr1 := []int{1, 2, 4, 5, 6}
	arr2 := []int{2, 3, 5, 7}
	union := unionOfSortedArrays(arr1, arr2)
	union2 := unionOfSortedArrays2(arr1, arr2)
	fmt.Println(union) // [1 2 3 4 5 6 7]
	fmt.Println(union2) // [1 2 3 4 5 6 7] (order may vary due to map iteration)
}
