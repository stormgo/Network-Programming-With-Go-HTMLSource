/* EchoClient
 */
package main

import (
	"fmt"
	"old/netchan"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	importer, err := netchan.Import("tcp", service)
	checkError(err)

	fmt.Println("Got importer")
	echoIn := make(chan int)
	importer.Import("echo-in", echoIn, netchan.Recv, 1)
	fmt.Println("Imported in")

	echoOut := make(chan int)
	importer.Import("echo-out", echoOut, netchan.Send, 1)
	fmt.Println("Imported out")

	for n := 1; n < 10; n++ {
		/*
			if closed(echoIn) {
				fmt.Println("In closed")
			}
		*/
		s := <-echoIn
		fmt.Println(s, n)
		echoOut <- s
	}
	// close(echoOut)
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
