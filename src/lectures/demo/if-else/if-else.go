package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("quiz 1 higher")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz 2 higher")
	} else {
		fmt.Println("Scores are equal")
	}

}
