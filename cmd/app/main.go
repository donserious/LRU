package main

import (
	"fmt"
	internal "lru/internal/cache"
)

func main() {
	cache := internal.NewLRUCache(4)
	cache.Add("1", "kek")
	cache.Add("2", "kekkek2")
	cache.Add("3", "kekkek3")
	cache.Add("4", "kekkek4")
	cache.Add("5", "kekkek5")

	cache.Add("3", "kekkekkek")
	fmt.Println(cache.Get("1"))
}
