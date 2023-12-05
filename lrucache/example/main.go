package main

import (
	"fmt"
	"time"

	"github.com/devetek/go-core/lrucache"
)

func main() {
	cachePerson := lrucache.New(5, time.Second*100)
	cacheAttributes := lrucache.New(100, time.Second*100)

	cachePerson.Set("name", "Nedya Prakasa")
	cacheAttributes.Set("posts", map[string]interface{}{
		"color":      "#fff",
		"background": "#000",
	})

	fmt.Println(cachePerson.Get("name"))
	fmt.Println(cacheAttributes.Get("posts"))
}
