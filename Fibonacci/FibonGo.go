package main

import (
	"fmt"
	"encoding/json"
)

var (n int)


func main() {

	index := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13"}
	value := []interface{}{0,1,1,2,3,5,8,13,21,34,55,89,144}

	m := make(map[string]interface{})

	cor := 0
	incor := 0

	for i, v  :=  range value {

		if cor != 10 && incor != 3 {

			fmt.Println("Ввведите число: ")
			fmt.Scanf("%d", &n)

			if n == v {
				fmt.Println("Welldone!")
				cor += 1

			} else {
				m[index[i]] = v
				jsonString, _ := json.Marshal(m)
				fmt.Printf("Wrong! \nCorrect is:  %s \n", jsonString)
				incor += 1

			}
			fmt.Println("Correct", cor)
			fmt.Println("Incorrect", incor)
			fmt.Println("\n")
		}
	}
}
