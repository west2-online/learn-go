package main

import (
	"fmt"
	"os"
)
func main(){
	var num int
	fmt.Scan(&num)
	if num < 17 || num > 1000000 {
		os.Exit(1)
	}else{
		result := isPrime(num)
		if result == true {
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
	}
}
func isPrime(x int) bool{
	for i := 2 ; i < x ; i++{
		if x % i == 0 {
			return false
		} 
	}
	return true
}