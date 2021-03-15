package main

import (
	"cache"
	"fmt"
	"time"
)

func main() {

	c := cache.New(5*time.Minute, 10*time.Minute)

	//c.Set("foo", 42, cache.NoExpiration)

	foo, found := c.Get("foo", nil)
	if found {
		fmt.Println(foo)
	}
}
