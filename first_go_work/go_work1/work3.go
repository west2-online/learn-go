package main

import "fmt"

func main() {
	var minyear, maxyear int
	count := 0
	slice := make([]int, 0)
	fmt.Scanln(&minyear, &maxyear)
	for i := minyear; i <= maxyear; i++ {
		if (i%4 == 0 && i%100 != 0) || (i%400 == 0) {
			count++
			slice = append(slice, i)
		}
	}
	fmt.Println(count)
	for i := range slice[:] {
		fmt.Print(slice[i], " ")
	}
}
