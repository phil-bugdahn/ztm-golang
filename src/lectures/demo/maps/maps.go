package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1 // Add 1 to current value, in this case the default (0)
	shoppingList["eggs"] += 1 // 12 eggs
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println(shoppingList)

	fmt.Println("need", shoppingList["eggs"])

	cereal, found := shoppingList["cereal"] // not found, cereal will have default value
	if !found {
		fmt.Println("Nope")
	} else {
		fmt.Println("Yup", cereal)
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("Total items", totalItems)
}
