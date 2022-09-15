package searching

import "fmt"

// recursion is a programming technique using function or
//algorithm that calls itself one or more times until a specified condition is met at which time the rest of each repetition
// is processed from the last one called to the first.

//If the size of the array is zero then, return -1,
//representing that the element is not found. This can also be treated as the base condition of a recursion call
//Otherwise, check if the element at the current index in the array is equal to the key or not i.e, arr[size – 1] == key
//If equal, then return the index of the found key

func searchRecursive(arr []int, size int, key int) int {

	if key == 0 { // The key is zero return -1
		return -1
	}

	if arr[size-1] == key {
		//return index of found key
		return size - 1
	} else {
		ans := searchRecursive(arr, size-1, key)
		return ans
	}

}

func main() {

	searchValue := 100
	size := 9

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(arr)

	ans := searchRecursive(arr, size, searchValue)
	if ans == -1 {
		fmt.Println("Key not found")
	} else {
		fmt.Printf("Key found at position: %d\n", ans)
	}

}
