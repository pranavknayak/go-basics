package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 { // Checks for proper usage
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]                               // service will be the string of the tcp address
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service) // converts it into a tcpAddr object containing an IP and a port number
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr) // establishes a connection to the remote server pointed to by tcpAddr, returning a connection type variable
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n")) // writes an HTTP request through the connection to the server
	checkError(err)
	result, err := ioutil.ReadAll(conn) //reads all the content from conn, which contains the response by the server
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) { // prints out the standard error message corresponding to the error passed as an argument
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
