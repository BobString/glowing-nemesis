package main

import (
	"encoding/gob"
	"failureDetector"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	process = map[int]string{
		1: "111.111.11.1:1200",
		2: "111.111.11.1:1200",
		3: "111.111.11.1:1200",
		4: "111.111.11.1:1200",
		5: "111.111.11.1:1200",
	}
	delay int = 2
)

func main() {
	//Entry point of the application

	//Create the connection
	createServer()

	//Launch Failure Detector		
	var proc []int
	keys := make([]int, len(process))
	i := 0
	for k, _ := range process {
		keys[i] = k
		i++
	}

	//FIXME: Should take the instance?? yes
	go failureDetector.init(delay, keys)

}

func createServer() {
	fmt.Println("Starting server...")
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	go listen(tcpAddr)

}

func listen(tcp *TCPAddr) {
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		fmt.Println("Listening.......")
		conn, err := listener.Accept()
		if err != nil {

			continue

		}
		handleClient(conn)
		//FIXME: Close it?
		//TODO: Try to put fori n handle client, one connection permanent per node
		conn.Close() // we’re finished

	}

}

func handleClient(conn net.Conn) {
	var buf [512]byte
	n, err := conn.Read(buf)
	checkError(err)
	println("received ", n, " bytes of data =", string(buf))
	var res []string
	res = strings.Split(string(buf), "@")
	switch {
	case res[0] == "Suspect":

	case res[0] == "Restore":

	case res[0] == "HeartbeatReply":
		//Call the method of the PFD
	case res[0] == "HeartbeatRequest":

	}

	//FIXME: reply it?
	////send reply
	//_, err = conn.Write(buf)
	//if err != nil {
	//	println("Error send reply:", err.Error())
	//} else {
	//	println("Reply sent")
	//}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
