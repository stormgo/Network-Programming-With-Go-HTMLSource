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
	err := exporter.ListenAndServe("tcp", ":2346")
	checkError(err)

	exportSend(exporter)

	importer, err := netchan.Import("tcp", "localhost:2346")
	checkError(err)
	importReceive(importer)

}
func exportSend(exp *netchan.Exporter) {
	ch := make(chan string)
	err := exp.Export("exportedSend", ch, netchan.Send)
	checkError(err)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- "hello"
		}
		close(ch)
		//exp.Drain(1000000)
	}()
}

func importReceive(imp *netchan.Importer) {
	ch := make(chan string)
	err := imp.ImportNValues("exportedSend", ch, netchan.Recv, 10, 1)
	checkError(err)

	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Println("Received \"" + v + "\"")
		/*
		                if closed(ch) {
					fmt.Println("Got close")
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
