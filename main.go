package main

import (
	"fmt"
	"log"

	"github.com/spankie/dchat/server"
)

var (
	id   uint
	name = "DEE CHAT"
)

func init() {
	log.SetPrefix(fmt.Sprintf("[%s]: ", name))
}

func main() {
	log.Printf("Start %s server\n", name)

	server := server.New()
	c := server.Start()

	for _ = range c {
		log.Println("Server going to sleep now...")
	}
}
