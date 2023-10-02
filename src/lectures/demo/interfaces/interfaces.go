package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cook chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("Chop salad")
	fmt.Println("add dressing")
}

func prepareDish(dishes []Preparer) {
	fmt.Println("PRepare dishes")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--dish: %v--\n", dish)
		dish.PrepareDish()
	}

	fmt.Println("")
}


func main() {
	dishes := []Preparer{Chicken("Chicken wings"), Salad("Iceberg salad")}
	prepareDish(dishes)
}
