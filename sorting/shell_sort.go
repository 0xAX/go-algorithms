package main

/*
 * Shell sort - http://en.wikipedia.org/wiki/Shellsort
 */

import "fmt"

import "github.com/0xAX/go-algorithms"

func main() {
	arr := utils.RandArray(10)
    fmt.Println("Initial array is:", arr)
    fmt.Println("")

    for d := int(len(arr)/2); d > 0; d /= 2 {
    	for i := d; i < len(arr); i++ {
    		for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
    			arr[j], arr[j-d] = arr[j-d], arr[j]
    		}
    	}
    }

    fmt.Println("Sorted array is: ", arr) 	
}