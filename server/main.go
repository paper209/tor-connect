package main

import (
	"server/admin"
	"server/client"
)

func main() {
	go client.Listen()
	admin.Handler()
}
