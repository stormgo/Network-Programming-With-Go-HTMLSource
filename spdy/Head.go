/* Head
 */
package main

import (
	"fmt"
	"http/spdy"
	"net"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: ", os.Args[0], " host:port\n")
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError("Dial", err)

	framer, err := spdy.NewFramer(conn, conn)
	checkError("Framer", err)

	// The client initiates a request by sending a SYN_STREAM frame.
	syn := &spdy.SynStreamFrame{
		/*
			CFHeader: spdy.ControlFrameHeader{
				Flags: spdy.ControlFlagFin,
				//version:   spdy.Version,
				//frameType: spdy.TypeSynStream,
			},
		*/
		Headers: http.Header{
			"Url":     []string{"http://www.google.com/"},
			"Method":  []string{"get"},
			"Version": []string{"http/1.1"},
		},
	}
	err = framer.WriteFrame(syn)
	checkError("Syn", err)

	/*
		for n := 0; n < 10; n++ {
			conn.Write([]byte("Hello " + string(n+48)))

			var buf [512]byte
			n, err := conn.Read(buf[0:])
			checkError("Read", err)

			fmt.Println(string(buf[0:n]))
		}
	*/
	os.Exit(0)
}

func checkError(errStr string, err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, errStr, err.Error())
		os.Exit(1)
	}
}
