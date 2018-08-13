package main
import "fmt"

func search(arr []int, key int) int {
  for i := 0; i < len(arr); i++ {
    if arr[i] == key {
      return i
    }
  }
  return -1
}


func main() {

  searchValue := 100
  arr := []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
  fmt.Println(arr)

  found := search(arr, searchValue)

  if found == -1 {
    fmt.Println("Key not found")
  } else {
    fmt.Printf("Key found at position: %d\n", found)
  }

}
