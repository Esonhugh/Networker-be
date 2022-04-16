package main

import (
	"Network-be/Server"
	"Network-be/config"
)

func main() {
	config.Init()
	Server.Run()

}
