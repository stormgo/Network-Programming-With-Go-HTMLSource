// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

func newServerConn(rwc io.ReadWriteCloser, buf *bufio.ReadWriter, req *http.Request,
	protocolChooser ProtocolChooser) (conn *Conn, err error) {
	config := new(Config)
	var hs serverHandshaker = &hybiServerHandshaker{Config: config}
	code, err := hs.ReadHandshake(buf.Reader, req, protocolChooser)
	if err == ErrBadWebSocketVersion {
		fmt.Fprintf(buf, "HTTP/1.1 %03d %s\r\n", code, http.StatusText(code))
		fmt.Fprintf(buf, "Sec-WebSocket-Version: %s\r\n", SupportedProtocolVersion)
		buf.WriteString("\r\n")
		buf.WriteString(err.Error())
		buf.Flush()
		return
	}
	if err != nil {
		hs = &hixie76ServerHandshaker{Config: config}
		code, err = hs.ReadHandshake(buf.Reader, req, protocolChooser)
	}
	if err != nil {
		hs = &hixie75ServerHandshaker{Config: config}
		code, err = hs.ReadHandshake(buf.Reader, req, protocolChooser)
	}
	if err != nil {
		fmt.Fprintf(buf, "HTTP/1.1 %03d %s\r\n", code, http.StatusText(code))
		buf.WriteString("\r\n")
		buf.WriteString(err.Error())
		buf.Flush()
		return
	}

	err = hs.AcceptHandshake(buf.Writer)
	if err != nil {
		code = http.StatusBadRequest
		fmt.Fprintf(buf, "HTTP/1.1 %03d %s\r\n", code, http.StatusText(code))
		buf.WriteString("\r\n")
		buf.Flush()
		return
	}
	conn = hs.NewServerConn(buf, rwc, req)
	return
}

/*
Handler is an interface to a WebSocket. It is a structure
with two fields: the Conn field is the function to handle 
messages between this server and user agents. The ProtocolChooser
field is for more advanced cases where the user agent and server
need to negotiate a subprotocol. It can usually be ignored

A trivial example server:

	package main

	import (
		"http"
		"io"
		"websocket"
	)

	// Echo the data received on the WebSocket.
	func EchoServer(ws *websocket.Conn) {
		io.Copy(ws, ws);
	}

	func main() {
		http.Handle("/echo", websocket.Handler(EchoServer));
		err := http.ListenAndServe(":12345", nil);
		if err != nil {
			panic("ListenAndServe: " + err.String())
		}
	}
*/
type HandlerStruct struct {
	conn            func(*Conn)
	ProtocolChooser ProtocolChooser
}

func Handler(h func(*Conn)) *HandlerStruct {
	return &HandlerStruct{conn: h}
}

// ServeHTTP implements the http.Handler interface for a Web Socket
func (h *HandlerStruct) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	rwc, buf, err := w.(http.Hijacker).Hijack()
	if err != nil {
		panic("Hijack failed: " + err.Error())
		return
	}
	// The server should abort the WebSocket connection if it finds
	// the client did not send a handshake that matches with protocol
	// specification.
	defer rwc.Close()
	conn, err := newServerConn(rwc, buf, req, h.ProtocolChooser)
	if err != nil {
		return
	}
	if conn == nil {
		panic("unexpected nil conn")
	}
	h.conn(conn)
}
