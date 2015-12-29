/* EchoServer
 */
package main

import (
	"fmt"
	"os"
	"old/netchan"
)

func main() {

	exporter := netchan.NewExporter()
	err := exporter.ListenAndServe("tcp", ":2345")
	checkError(err)

	echoIn := make(chan int)
	echoOut := make(chan int)
	exporter.Export("echo-in", echoIn, netchan.Send)
	exporter.Export("echo-out", echoOut, netchan.Recv)
	for s := 0; ; s++ {
		/*
			if closed(echoIn) {
				fmt.Println("Channel closed")
				break
			}
		*/
		fmt.Println("Sending to echoIn")
		echoIn <- s
		fmt.Println("Sent to echoIn")
		/*
			 if closed(echoOut) {
				fmt.Println("Channel closed")
				break
			}
		*/
		fmt.Println("Getting from echoOut")
		r := <-echoOut
		fmt.Println("Got from echoOut")
		/*
			if closed(echoOut) {
				fmt.Println("Channel closed")
				break
			}
		*/
		fmt.Println("received", r)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
