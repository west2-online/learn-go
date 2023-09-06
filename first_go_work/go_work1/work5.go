package main

import "fmt"

func main() {
	slice1 := make([]int, 0)
	for i := 1; i < 51; i++ {
		slice1 = append(slice1, i)
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i]%3 == 0 {
			slice1 = append(slice1[:i], slice1[i+1:]...)
			i--
		}
	}
	slice1 = append(slice1, 114514)
	fmt.Println(slice1)
}
