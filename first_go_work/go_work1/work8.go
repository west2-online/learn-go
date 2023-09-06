package main

import "fmt"

func main() {
	var num = map[int]int{
		2:  1,
		7:  2,
		11: 3,
		15: 4,
	}
	var target int
	fmt.Scanf("%d", &target)
	for k, v := range num {
		val, ok := num[target-k]
		if ok {
			fmt.Println(v, val)
			delete(num, k)
		}
	}
}
