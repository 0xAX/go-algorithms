package main

/*
 * Comb sort - https://en.wikipedia.org/wiki/Combsort
 */

import "fmt"

import "github.com/0xAX/go-algorithms"

func main() {
    arr := utils.RandArray(10)
    fmt.Println("Initial array is:", arr)
    fmt.Println("")
    
    tmp := 0
    gap := len(arr)
    
    for {
        if gap > 1 {
            gap = gap * 100 / 124
        }
        
        for i := 0 ; ; {
        
            if arr[i] > arr[i + gap]  {         
                tmp = arr[i]
                arr[i] = arr[i + gap]
                arr[i + gap] = tmp              
            }

            i++
            
            if i + gap >= len(arr){
                break
            }
        }

        if gap == 1 {
            break
        }
    }

    fmt.Println("Sorted array is: ", arr)
}
