package main

import (
	"fmt"
	"os"
	"sort"
)
func main(){
	var year1 int 
	var year2 int
	var num1 int = 0
	var num2 int = 0
	var sum int
	var slice []int
	fmt.Scan(&year1,&year2)
	if year1 >=1582 && year2 >= year1 &&year2<=3000 {
		for a := year1 ; a <= year2 ; a++ {
			if a%100 == 0{
				if a%400 ==0{
					slice = append(slice, a)
					num1++
				}
			}else{
				if a%4 ==0{
					slice = append(slice, a)
					num2++
				}
			}
		}
		sum = num1 + num2
		fmt.Println(sum)
		sort.Ints(slice)
		for i, c := range slice {
			fmt.Print(c)
			if i < len(slice)-1 {
				fmt.Print(" ")
			}
		}
	}else{
		os.Exit(1)
	}
}