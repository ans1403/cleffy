package main

import (
	"cleffy/src/router"
)

func main() {
	r := router.GetRouter()
	r.Run(":8080")
}
