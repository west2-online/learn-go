package main

import "fmt"

func main() {
	var arr [10]int
	var min int
	count := 0

	fmt.Scanln(&arr[0], &arr[1], &arr[2], &arr[3], &arr[4], &arr[5], &arr[6], &arr[7], &arr[8], &arr[9])

	fmt.Scanf("%d", &min)
	for i := 0; i < 10; i++ {
		if arr[i] <= (min + 30) {
			count++
		}
	}
	fmt.Println(count)
}
