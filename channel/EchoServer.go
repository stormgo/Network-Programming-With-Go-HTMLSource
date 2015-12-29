/* EchoServer
 */
package main

import (
	"fmt"
	"os"
	"old/netchan"
)

func main() {

	// exporter, err := netchan.NewExporter("tcp", ":2345")
	exporter := netchan.NewExporter()
	err := exporter.ListenAndServe("tcp", ":2345")
	checkError(err)

	echoIn := make(chan string)
	echoOut := make(chan string)
	exporter.Export("echo-in", echoIn, netchan.Send)
	exporter.Export("echo-out", echoOut, netchan.Recv)
	for {
		fmt.Println("Getting from echoOut")
		s, ok := <-echoOut
		if !ok {
			fmt.Printf("Read from channel failed")
			os.Exit(1)
		}
		fmt.Println("received", s)

		fmt.Println("Sending back to echoIn")
		echoIn <- s
		fmt.Println("Sent to echoIn")
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
