package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	nineFile, err := os.OpenFile("./ninenine.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
	}
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Fprint(nineFile, i, " * ", j, " = ", i*j, " ")
		}
		fmt.Fprint(nineFile, "\n")
	}

}
