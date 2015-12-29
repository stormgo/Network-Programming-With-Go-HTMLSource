package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"reflect"
	"websocket"
)

var XML = websocket.Codec{xmlMarshal, xmlUnmarshal}

func xmlMarshal(v interface{}) (msg []byte, payloadType byte, err error) {
	buff := &bytes.Buffer{}
	Marshal(v, buff)
	msg = buff.Bytes()
	return msg, websocket.TextFrame, nil
}

func xmlUnmarshal(msg []byte, payloadType byte, v interface{}) (err error) {
	r := bytes.NewBuffer(msg)
	err = xml.Unmarshal(r, v)
	return err
}

func Marshal(e interface{}, w io.Writer) {
	// make it a legal XML document
	w.Write([]byte("<?xml version=\"1.1\" encoding=\"UTF-8\" ?>\n"))

	// topvel e is a value and has no structure field, 
	// so use its type
	typ := reflect.TypeOf(e)
	name := typ.Name()

	startTag(name, w)
	MarshalValue(reflect.ValueOf(e), w)
	endTag(name, w)
}

func MarshalValue(v reflect.Value, w io.Writer) {
	t := v.Type()
	switch t.Kind() {
	case reflect.Struct:
		for n := 0; n < t.NumField(); n++ {
			field := t.Field(n)

			vv := v

			// special case if it is a slice

			if vv.Field(n).Type().Kind() == reflect.Slice {
				// slice
				MarshalSliceValue(field.Name,
					vv.Field(n), w)
			} else {
				// not a slice
				startTag(field.Name, w)
				MarshalValue(vv.Field(n), w)
				endTag(field.Name, w)
			}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
	case reflect.Bool:
	case reflect.String:
		vv := v
		w.Write([]byte("   " + vv.String() + "\n"))
	default:
	}
}

func MarshalSliceValue(tag string, v reflect.Value, w io.Writer) {
	for n := 0; n < v.Len(); n++ {
		startTag(tag, w)
		MarshalValue(v.Index(n), w)
		endTag(tag, w)
	}
}

func startTag(s string, w io.Writer) {
	w.Write([]byte("<" + s + ">\n"))
}

func endTag(s string, w io.Writer) {
	w.Write([]byte("</" + s + ">\n"))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
