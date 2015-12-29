package main

import (
	"dictionary"
	"fmt"
)

func main() {
	dictionaryPath := "/var/www/html/go/chinese/cedict_ts.u8"
	d := new(dictionary.Dictionary)
	d.Load(dictionaryPath)

	haoD := d.LookupPinyin("hao2")
	fmt.Println(haoD.String())

	goodD := d.LookupEnglish("good")
	fmt.Println(goodD.String())
}
