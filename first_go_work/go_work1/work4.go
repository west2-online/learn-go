package main

import "fmt"

func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var a int
	fmt.Scanf("%d", &a)
	if isPrime(a) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
