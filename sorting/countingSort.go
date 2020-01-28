package main

import "fmt"

//use this for strings
func countingSortString(data []rune, size int) string {
	length := 256
	count := make([]int, length)
	output := make([]rune, size)
	for i := 0; i < size; i++ {
		count[data[i]]++
	}

	for i := 1; i < length; i++ {
		count[i] += count[i-1]
	}

	for i := 0; i < size; i++ {
		output[count[data[i]]-1] = data[i]
		count[data[i]]--
	}

	return string(output)
}

func countingSortNumber(data []int, size int) []int {
	min := data[0]
	for _, val := range data {
		if val < min {
			min = val
		}
	}
	max := data[0]
	for _, val := range data {
		if val > max {
			max = val
		}
	}

	length := max - min + 1
	count := make([]int, length)
	output := make([]int, size)
	for i := 0; i < size; i++ {
		count[data[i]-min]++
	}

	for i := 1; i < length; i++ {
		count[i] += count[i-1]
	}

	for i := 0; i < size; i++ {
		output[count[data[i]-min]-1] = data[i]
		count[data[i]-min]--
	}

	return output

}

func main() {
	data := "golang"
	g := []rune(data)
	str := countingSortString(g, len(data))
	fmt.Println(str)

	arr := []int{0, -2, -9, 20, 7}
	fmt.Println(countingSortNumber(arr, len(arr)))
}
