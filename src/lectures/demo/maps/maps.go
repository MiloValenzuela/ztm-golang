package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)

	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)
	fmt.Println("Milk deleted, new list:", shoppingList)

	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal")
	if !found {
		fmt.Println("nope")
	} else {
		fmt.Println("yep", cereal, "boxes")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("There are", totalItems, "on the shopping list")
}
