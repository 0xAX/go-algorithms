package main

/*
 * Merge sort - http://en.wikipedia.org/wiki/Merge_sort
 */

func Merge(left, right []int) []int {
    result := make([]int, 0, len(left) + len(right))
    
    for len(left) > 0 || len(right) > 0 {
        if len(left) == 0 {
            return append(result, right...)
        }
        if len(right) == 0 {
            return append(result, left...)
        }
        if left[0] <= right[0] {
            result = append(result, left[0])
            left = left[1:]
        } else {
            result = append(result, right[0])
            right = right[1:]
        }
    }
    
    return result
}

func MergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    
    middle := len(arr) / 2
    
    left := MergeSort(arr[:middle])
    right := MergeSort(arr[middle:])
    
    return Merge(left, right)
}
