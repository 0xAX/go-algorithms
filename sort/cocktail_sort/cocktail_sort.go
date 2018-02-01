package main

/*
 * Cocktail sort - https://en.wikipedia.org/wiki/Cocktail_sort
 */

import "fmt"

import "github.com/0xAX/go-algorithms"

func main() {
    arr := utils.RandArray(10)
    fmt.Println("Initial array is:", arr)
    fmt.Println("")
    
    tmp   := 0
    
    for i := 0; i < len(arr) / 2; i++ {
        left := 0
        right := len(arr) - 1
        
        for ; left <= right ; {
            
            if arr[left] > arr[left + 1] {
                tmp = arr[left]
                arr[left] = arr[left + 1]
                arr[left + 1] = tmp
            }
			
            left++
            
            if arr[right - 1] > arr[right] {
                tmp = arr[right - 1]
                arr[right - 1] = arr[right]
                arr[right] = tmp
            }

            right--
        }
    }
    
    fmt.Println("Sorted array is: ", arr)
}
