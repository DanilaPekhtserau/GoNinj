package main

import (
	"fmt"
	"golang-ninj/cache"
)

func main() {
	c := cache.New()

	c.Set("userId", 42)
	userId := c.Get("userId")

	fmt.Println(userId)

	c.Delete("userId")
	userId = c.Get("userId")

	fmt.Println(userId)
}
