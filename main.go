package main

import "github.com/kinokoruumu/vue-go-blog/router"

func main() {
	r := router.GetRouter()
	r.Run(":2000")
}
