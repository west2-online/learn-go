package main
import("fmt")
func main(){
	var slice []int
	for i := 1 ; i <= 50 ; i++ {
		slice = append(slice, i)
	}
	for a := 0 ; a < len(slice) ; a++ {
		if slice[a]%3 == 0 {
			indexToRemove := a
			slice = append(slice[:indexToRemove],slice[indexToRemove+1:]...)
		}
	}
	b := 114514
	slice = append(slice, b)
	fmt.Println(slice)
}