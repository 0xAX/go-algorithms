package main

/*
 * Selection sort - http://en.wikipedia.org/wiki/Selection_sort
 */

import "fmt"

import "github.com/0xAX/go-algorithms"

func main() {
    arr := utils.RandArray(10)
    fmt.Println("Initial array is:", arr)
    fmt.Println("")

    var min int = 0
    var tmp int = 0

    for i := 0; i < len(arr); i++ {
        min = i
        for j := i + 1; j < len(arr); j++ {
            if arr[j] < arr[min] {
                min = j
            }
        }
		
        tmp = arr[i]
        arr[i] = arr[min]
        arr[min] = tmp
    }

    fmt.Println("Sorted array:    ", arr)
}
