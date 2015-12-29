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

}
func exportSend(exp *netchan.Exporter) {
	ch := make(chan string)
	err := exp.Export("exportedSend", ch, netchan.Send)
	checkError(err)

	func() {
		for i := 0; i < 5; i++ {
			ch <- "hello"
		}
		close(ch)
		//exp.Drain(1000000)
	}()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
