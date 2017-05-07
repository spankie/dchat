package server

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	// defaultport defines port the server listens on
	defaultport = "8080"
)

var (
	IDs int
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

			log.Printf("connection %d accepted\n", clientID)
			conn.Write([]byte(strconv.Itoa(clientID)))
			go read(conn, clientID)
			go write(conn, clientID)
		}
		channel <- 1
	}()

	return channel
	// TODO:: Launch a go routine to handle incoming requests...
}

func read(c net.Conn, id int) {
	bb := make([]byte, 8)
	for {
		// copy 8Bytes at a time
		// _, err := io.Copy(os.Stdout, c)
		_, err := c.Read(bb)
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				log.Println("Connection is closed.")
				break
			}
			log.Println("could not copy message from client")
			break
		}
		fmt.Printf("[%d] : %s\n", id, string(bb))
		// log.Println("\nHi: ")
	}
}

func write(c net.Conn, id int) {
	for {
		c.Write([]byte(fmt.Sprintf("\nServer: Welcome client: %#v\n", id)))
		time.Sleep(time.Second * 3)
		// break
	}
}
