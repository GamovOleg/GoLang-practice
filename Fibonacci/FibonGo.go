package main

import (
	"fmt"
)

var (n int)

func main() {

	fib := [10]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34,}
	cor :=0
	incor :=0

	for index, value  :=  range fib{
		if cor !=10 && incor !=3{

		fmt.Println("Ввведите число: ")
		fmt.Scanf("%d", &n)
		if n == value {
			fmt.Println("Welldone!")
			cor += 1
		} else {
				fmt.Println("Wrong! \n", index, value)
				incor += 1
			}
			fmt.Println("Correct", cor)
			fmt.Println("Incorrect", incor)
			fmt.Println("\n")
			}

	}
}


