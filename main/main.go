package main

import (
	router "userapi/router"
)

func main() {
	r := router.NewRoute()
	r.Run(":8080")

}
