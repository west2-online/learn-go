package main
import (
	"math"
	"fmt"
	"os"
)	
func main(){
	var num1,num2 int
	fmt.Scan(&num1,&num2)
	absNum1 := math.Abs(float64(num1))
	absNum2 := math.Abs(float64(num2))
	var sum = 0;
	if absNum1<=1e9 && absNum2<=1e9 {
		sum = num1 + num2
		fmt.Println(sum)
	}else{
		os.Exit(1)
	}						
}