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
	importSend(importer)
	fmt.Println("Received!")
	os.Exit(0)
}

func importSend(imp *netchan.Importer) {
	ch := make(chan int)
	err := imp.ImportNValues("exportedRecv", ch, netchan.Send, count, 1)
	if err != nil {
		fmt.Println("importSend:", err)
	}
	for i := 0; i < closeCount; i++ {
		ch <- 45 + i
	}
	close(ch)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
