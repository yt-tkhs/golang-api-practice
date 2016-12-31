package main

import "main/route"

func main() {
	router := route.InitRouter()
	router.Start(":1323")
}