/* EchoClient
 */
package main

import (
	"fmt"
	"old/netchan"
	"os"
)

const count = 10     // number of items in most tests
const closeCount = 5 // number of items when sender closes early


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	importer, err := netchan.Import("tcp", service)
	checkError(err)
	importReceive(importer)
	fmt.Println("Received!")
	os.Exit(0)
}

func importReceive(imp *netchan.Importer) {
	ch := make(chan string)
	err := imp.ImportNValues("exportedSend", ch, netchan.Recv, count, 1)
	checkError(err)

	for i := 0; i < count; i++ {
		v := <-ch
		fmt.Println("Received \"" + v + "\"")
		/*
		                if closed(ch) {
					fmt.Println("Got close")
		                        if i != closeCount {
		                                fmt.Println("expected close at %d; got one at %d\n", closeCount, i)
		                        }
		                        break
		                }
		*/
		/*
		   if v != 23+i {
		           fmt.Println("importReceive: bad value: expected %d; got %+d", 23+i, v)
		   }
		*/
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
