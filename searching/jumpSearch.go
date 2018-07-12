package main
import "fmt"
import "math"

func jumpSearch(arr []int, key int) int {

  //block size to jump
  sz := len(arr)
  step := int(math.Sqrt(float64(sz)))
  prev := 0

  //finding the block
  for arr[int(math.Min(float64(step), float64(sz))) - 1] < key {
    prev = step
    step += int(math.Sqrt(float64(sz)))
    if prev >= sz {
      return -1
    }
  }

  //linear search the block
  for arr[prev] < key {
    prev++
    if prev == int(math.Min(float64(step), float64(sz))) {
      return -1
    }
  }

  if arr[prev] == key {
    return prev
  }
  return -1
}


func main() {

  arr := []int { 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610 }
  key := 55

  index := jumpSearch(arr, key)
  fmt.Println(index)

}
