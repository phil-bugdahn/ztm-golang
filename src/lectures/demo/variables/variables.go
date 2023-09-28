package main

import "fmt"

func main() {
	var myName = "Phil"
	fmt.Println("my name is ", myName)

	var name string = "Kathy"
	fmt.Println("name = ", name)

	username := "admin"
	fmt.Println("username = ", username)

	var sum int
	fmt.Println("The sum is ", sum)

	part1, other := 1, 5
	fmt.Println("part 1 is ", part1, " other is ", other)

	part2, other := 1, 0
	fmt.Println("part 1 is ", part2, " other is ", other)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName = ", lessonName)
	fmt.Println("lessonType = ", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println("word 1 is ", word1, " word2 is ", word2)
}
