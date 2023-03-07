package main

import "fmt"

func ShiftArray(array []int, i int) []int {
	for x := 1; x <= i; x++ {
		tmp := []int{
			array[3], array[0], array[1],
			array[6], array[4], array[2],
			array[7], array[8], array[5],
		}
		array = tmp
	}
	return array
}

func main() {
	fmt.Println(ShiftArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1))
}
