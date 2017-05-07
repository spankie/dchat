package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

const (
	// defaultport defines port the server listens on
	defaultport = "8080"
)

var (
	// IDs hold number for assigning IDs to clients
	IDs         int
	messageChan = make(chan string)
	clientsConn []net.Conn
)

// Server is a structure for the server
type Server struct {
	ServePort string
	Clients   int
}

// New creates a default server with default values
func New() *Server {
	return &Server{ServePort: defaultport}
}

// Start starts the server and returns a listener
func (s *Server) Start() chan int {
	log.Println("Server started on ", s.ServePort)

	// Start the server ...
	server, err := net.Listen("tcp", fmt.Sprintf(":%s", s.ServePort))
	if err != nil {
		// end the program if this server could not be started
		log.Fatalln("Could not start server. reason: ", err)
	}

	// Control channel for the server
	channel := make(chan int)

	// TODO:: implement spreading clients message to other clients

	go func() {
		for {
			// Limit to four clients for now...
			if IDs == 4 {
				break
			}
			log.Println("Waiting for client...")
			conn, err := server.Accept()
			if err != nil {
				log.Println("error in connection.")
				continue
			}
			// add one to the number of connected clients
			s.Clients++
			// increase id for uniquesness
			IDs++
			// Assign
			clientID := IDs
			clientsConn = append(clientsConn, conn)
			log.Printf("connection %d accepted\n", clientID)
			conn.Write([]byte(strconv.Itoa(clientID) + "\n"))
			go read(conn, clientID)
		}
		channel <- 1
	}()
	go write()
	return channel
	// TODO:: Launch a go routine to handle incoming requests...
}

func read(c net.Conn, id int) {
	// check if the connection is closed with channels
	// bb := make([]byte, 8)
	buf := bufio.NewReader(c)
	for c != nil {

		message, _, err := buf.ReadLine()
		if err != nil {
			log.Println("Could not get message")
			break
		}
		m := string(message)
		m = fmt.Sprintf("[%d] > %s\n", id, m)
		// fmt.Printf("[%d] : %s\n", id, string(message))
		// fmt.Print(m)
		messageChan <- m
	}
	log.Printf("Connection[%d] Is closed.\n", id)
}

func write() {
	// check if the connection is closed with channels
	for message := range messageChan {
		for _, c := range clientsConn {
			// message = message + "\n"
			// log.Println("m:", message)
			c.Write([]byte(message))
			// time.Sleep(time.Second * 3)
		}
		// break
	}
}
