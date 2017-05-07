package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	myID string
)

// type Rat struct {
// 	Conn net.Conn
// }

// func (r Rat) ReadAt(p []byte, off int64) (n int, err error) {
// 	p = p[off : off+off]
// 	return int(off), nil
// }

func main() {
	// connect to server
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln("could not dial the server")
		return
	}

	// bw := make([]byte, 32)
	// io.ReadFull(os.Stdin, bw)
	// conn.Write(bw)

	// bb := make([]byte, 16)
	var bb = make([]byte, 32<<10)

	_, _ = conn.Read(bb)
	myID = string(bb)
	if err != nil {
		log.Println(string(bb))
		log.Fatalln("Could not get ID")
	}
	log.Println("My ID : ", myID)

	go func() {
		for {
			_, _ = conn.Read(bb)
			fmt.Print(string(bb))
			// _, _ = io.WriteString(os.Stdout, string(bb))
			// if err == io.ErrUnexpectedEOF {
			// 	_, err = io.WriteString(os.Stdout, string(bb))
			// }
		}
	}()

	// bw := make([]byte, 32)
	for {
		io.Copy(conn, os.Stdin)
		// conn.Write(bw)
	}
	// copy 8Bytes
	// _, err = io.CopyN(os.Stdout, conn, 8)
	// conn.Write([]byte("Hello dee chat...\n")) // send EOF signal to the connection
	// w, err := io.Copy(os.Stdout, conn)
	// if err != nil {
	// 	log.Println("error encountered reading from the connection")
	// 	return
	// }
	// log.Println("Number of bytes received: ", w)
}
