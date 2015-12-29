/* EchoServer
 */
package main

import (
	"fmt"
	"os"
	"old/netchan"
)

const count = 10     // number of items in most tests
const closeCount = 5 // number of items when sender closes early

func main() {

	exporter := netchan.NewExporter()
	err := exporter.ListenAndServe("tcp", ":2346")
	checkError(err)

	exportReceive(exporter)

}

func exportReceive(exp *netchan.Exporter) {
	ch := make(chan int)
	err := exp.Export("exportedRecv", ch, netchan.Recv)
	checkError(err)
	for i := 0; ; i++ {
		/*
			if closed(ch) {
				fmt.Println("Channel closed")
			}
		*/
		v := <-ch
		fmt.Println("Received", v)
		/*
		   if v != 45+i {
		           fmt.Println("export Receive: bad value: expected 4%d; got %d", 45+i, v)
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
