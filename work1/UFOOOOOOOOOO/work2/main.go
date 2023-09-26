package main

import (
	"fmt"
	
)
func main(){
	var nums [10]int
	fmt.Scan(&nums[0],&nums[1],&nums[2],&nums[3],&nums[4],&nums[5],&nums[6],&nums[7],&nums[8],&nums[9])
	max := nums[0]
    min := nums[0]
	for _, num := range nums {
        if num > max {
            max = num
        }
        if num < min {
            min = num
        }
    }
	var len int
	fmt.Scan(&len)
	len = len +30
	var b int
	if min>=100&&max<=200{
		for a := 0; a <=9; a++{
				
			if len >= nums[a]{
				b++
			}else{
				continue
			}
				
		}
		fmt.Println(b)
	}
}
	
