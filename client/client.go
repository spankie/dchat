package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	myID string
)

func main() {
	// connect to server
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln("could not dial the server")
		return
	}

	var bb = make([]byte, 32<<10)
	// 32<<10 means multiply 32 by 2, 10 times.
	
	_, _ = conn.Read(bb)
	myID = string(bb)
	if err != nil {
		log.Println(string(bb))
		log.Fatalln("Could not get ID")
	}
	log.Println("My ID : ", myID)

	go func() {
		buf := bufio.NewReader(conn)
		for {
			// _, _ = conn.Read(bb)
			message, err := buf.ReadBytes('\n')
			if err != nil {
				log.Fatalln("Server Down! GoodBye!")
				break
			}
			fmt.Print(string(message))
		}
	}()

	buf := bufio.NewReader(os.Stdin)
	for {
		message, err := buf.ReadBytes('\n')
		if err != nil {
			log.Println("Could not get message.")
		}
		// fmt.Println(message)
		// send message to the server
		conn.Write(message)
	}
}
