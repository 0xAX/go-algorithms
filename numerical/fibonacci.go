package main
import "fmt"

//using recursion
func fibo(num int) int {
  if num <= 1 {
    return num
  }
  return fibo(num -1) + fibo(num - 2)
}

func main(){
  num := 10
  result := fibo(num)
  fmt.Println(result)
}
